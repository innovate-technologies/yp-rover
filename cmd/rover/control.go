package main

import (
	"log"

	"github.com/innovate-technologies/yp-rover/internal/cron"
	"github.com/spf13/cobra"
)

// controlCmd represents the control command
var controlCmd = &cobra.Command{
	Use:   "control",
	Short: "This starts a controller server instance",
	RunE:  runControl,
}

func init() {
	RootCmd.AddCommand(controlCmd)
}

func runControl(cmd *cobra.Command, args []string) error {
	ch, q, err := getQueue()
	if err != nil {
		return err
	}

	updateGenres := cron.UpdateGenres()
	forever := make(chan bool)

	go func() {
		for {
			select {
			case task := <-updateGenres:
				queueTask(ch, q, task)
			}
		}
	}()

	log.Printf(" [*] Running controller. To exit press CTRL+C")
	<-forever

	return nil
}
