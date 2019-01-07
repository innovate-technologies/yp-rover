package main

import (
	"errors"

	"github.com/innovate-technologies/yp-rover/internal/tasks"
	"github.com/spf13/cobra"
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "This launches a specific task",
	RunE:  runlaunch,
}

func init() {
	RootCmd.AddCommand(launchCmd)
}

func runlaunch(cmd *cobra.Command, args []string) error {
	ch, q, err := getQueue()
	if err != nil {
		return err
	}

	if len(args) < 2 {
		return errors.New("Needs 2 arguments: unit and function")
	}

	return queueTask(ch, q, tasks.Task{
		Unit:     args[0],
		Function: args[1],
	})
}
