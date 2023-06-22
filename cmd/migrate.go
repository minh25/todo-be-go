package cmd

import (
	"fmt"
	"todoProject/config"
	"todoProject/pkg/migration"
)

func init() {
	conf, err := config.Load()
	if err != nil {
		fmt.Println("Error:", err)
	}
	migrateCmd := migration.MigrateCommand(conf.MySQL.DSN())
	rootCmd.AddCommand(migrateCmd)
}
