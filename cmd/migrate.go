package cmd

import (
	"github.com/juancwu/go-ntt/util"
	"github.com/spf13/cobra"
)

const (
	migrationTable = "_gontt_migrations"
)

func migrate() *cobra.Command {
	var (
		source string
		conn   string
		prefix string
	)

	migrateCmd := &cobra.Command{
		Use:     "migrate",
		Short:   "Apply migrations",
		Long:    "Applies all 'up' migrations that haven't been done or are new",
		Example: "ntt migrate --source DIR_NAME --conn mysql://DB_CONN_STR [--prefix PREFIX]",
		RunE: func(cmd *cobra.Command, args []string) error {
			util.Log().Infof("Source: %s\n", source)
			util.Log().Infof("Conn: %s\n", conn)
			util.Log().Infof("Prefix: %s\n", prefix)
			return nil
		},
	}

	migrateCmd.Flags().StringVarP(&source, "source", "", "", sourceFlagUsage)
	migrateCmd.MarkFlagDirname("source")
	migrateCmd.MarkFlagRequired("source")

	migrateCmd.Flags().StringVarP(&conn, "conn", "", "", "This is a connection string that must include the database scheme/type.")
	migrateCmd.MarkFlagRequired("conn")

	migrateCmd.Flags().StringVarP(&prefix, "prefix", "", "", "This is a prefix for the migrations table. This allow using the same database for multiple migration tables for different projects.")

	return migrateCmd
}
