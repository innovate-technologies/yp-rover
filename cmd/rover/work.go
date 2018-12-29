package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// workCmd represents the work command
var workCmd = &cobra.Command{
	Use:   "work",
	Short: "This starts a worker server instance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("work called")
	},
}

func init() {
	RootCmd.AddCommand(workCmd)
}
