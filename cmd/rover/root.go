package main

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "rover",
	Short: "A lonely robot rolling its way on a far away planet looking for signs of music.",
	Long: `A lonely robot rolling its way on a far away planet looking for signs of music.

Which is a cryptic way to tell that this is a distributed discovery system for internet radio stations.
Which is a technical way of saying "The Google of Radio".
Which is an elevator pitch for discover.fm`,
}
