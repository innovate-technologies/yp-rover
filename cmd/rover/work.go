package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/innovate-technologies/yp-rover/internal/tasks/shoutcast"
	tunein "github.com/innovate-technologies/yp-rover/internal/tasks/tunein"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

// workCmd represents the work command
var workCmd = &cobra.Command{
	Use:   "work",
	Short: "This starts a worker server instance",
	RunE:  runWork,
}

func init() {
	RootCmd.AddCommand(workCmd)
}

func runWork(cmd *cobra.Command, args []string) error {
	ch, q, err := getQueue()
	if err != nil {
		return err
	}

	// set QOS to only give workers 1 task at a time
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
	L:
		for {
			select {
			case <-ctx.Done():
				log.Println("Got context break")
				break L
			case d := <-msgs:
				log.Println("Got message")
				if d.ContentType != "application/json" { // no idea who sent that
					d.Ack(false)
					continue
				}
				task := tasks.Task{}
				json.Unmarshal(d.Body, &task)
				log.Printf("Got task %s.%s\n", task.Unit, task.Function)

				handleTask(ch, q, task)
				d.Ack(false)
				log.Printf("End task %s.%s\n", task.Unit, task.Function)
				break
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-signals
	cancel()
	log.Printf("Got shutdown signal, finishing the work and exiting")

	return nil
}

func handleTask(ch *amqp.Channel, q *amqp.Queue, task tasks.Task) {
	var followUpTasks []tasks.Task
	var err error
	switch task.Unit {
	case "shoutcastcom":
		handler := shoutcast.New(config)
		followUpTasks, err = handler.HandleTask(task)
		break
	case "tunein":
		handler := tunein.New(config)
		followUpTasks, err = handler.HandleTask(task)
		break
	}

	if err != nil {
		log.Println(err)
		time.Sleep(time.Second)
		log.Printf("Got error, requeueing")
		queueTask(ch, q, task) // retry me
		return
	}
	if followUpTasks == nil {
		return
	}
	for _, newtask := range followUpTasks {
		err = queueTask(ch, q, newtask)
		if err != nil {
			log.Println(err)
		}
	}
}
