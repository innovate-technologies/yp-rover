package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// controlCmd represents the control command
var controlCmd = &cobra.Command{
	Use:   "control",
	Short: "This starts a controller server instance",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("control called")
	},
}

func init() {
	RootCmd.AddCommand(controlCmd)
}
