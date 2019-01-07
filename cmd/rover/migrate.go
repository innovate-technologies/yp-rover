package main

import (
	"github.com/innovate-technologies/yp-rover/pkg/store"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This starts a database table migration",
	RunE:  runControl,
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}

func runMigrate(cmd *cobra.Command, args []string) error {
	db, err := store.New(config.MySQLURL)
	if err != nil {
		return err
	}
	db.Migrate()

	return nil
}
