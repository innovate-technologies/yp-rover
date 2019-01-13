package main

import (
	"github.com/innovate-technologies/yp-rover/pkg/store"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This starts a database table migration",
	RunE:  runMigrate,
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}

func runMigrate(cmd *cobra.Command, args []string) error {
	db, err := store.New(config)
	if err != nil {
		return err
	}
	defer db.Close()
	db.Migrate()

	return nil
}
