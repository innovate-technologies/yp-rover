package main

import (
	"encoding/json"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/streadway/amqp"
)

func getQueue() (*amqp.Channel, *amqp.Queue, error) {
	conn, err := amqp.Dial(config.RabbitMQURL)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	q, err := ch.QueueDeclare(
		"yp-rover", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil, nil, err
	}

	return ch, &q, nil
}

func queueTask(ch *amqp.Channel, q *amqp.Queue, task tasks.Task) error {
	t, err := json.Marshal(task)
	if err != nil {
		return err
	}
	return ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         t,
		})
}
