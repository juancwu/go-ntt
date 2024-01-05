package cmd

import (
	"github.com/spf13/cobra"

	"github.com/juancwu/go-ntt/databases"
	"github.com/juancwu/go-ntt/util"
)

const (
	migrationTable = "_gontt_migrations"
)

func migrate() *cobra.Command {
	var (
		dir         string
		conn        string
		prefix      string
		downMigrate bool
	)

	migrateCmd := &cobra.Command{
		Use:     "migrate",
		Short:   "Apply migrations",
		Long:    "Applies all 'up' migrations that haven't been done or are new by default unless --down or -d is declared.",
		Example: "ntt migrate --dir DIR_NAME --conn SCHEME://CONN_STR [--prefix PREFIX] [--down]",
		RunE: func(cmd *cobra.Command, args []string) error {
			scheme, err := util.SchemeFromURL(conn)
			if err != nil {
				return err
			}

			d, err := databases.GetDriver(scheme)
			if err != nil {
				return err
			}

			util.Log().Info("opening database connection...")
			if err = d.Open(conn, prefix); err != nil {
				return err
			}

			if downMigrate {
				util.Log().Info("starting down migrations...")
				err = d.Down(dir)
			} else {
				util.Log().Info("starting up migrations...")
				err = d.Up(dir)
			}

			if err != nil {
				util.Log().Error("error running migrations.", "err", err)
				return err
			}

			return nil
		},
	}

	migrateCmd.Flags().StringVarP(&dir, "dir", "", "", dirFlagUsage)
	migrateCmd.MarkFlagDirname("dir")
	migrateCmd.MarkFlagRequired("dir")

	migrateCmd.Flags().StringVarP(&conn, "conn", "", "", "This is a connection string that must include the database scheme/type.")
	migrateCmd.MarkFlagRequired("conn")

	migrateCmd.Flags().StringVarP(&prefix, "prefix", "", "", "This is a prefix for the migrations table. This allow using the same database for multiple migration tables for different projects.")

	migrateCmd.Flags().BoolVarP(&downMigrate, "down", "d", false, "This will run down migrations, which will reset the database.")

	return migrateCmd
}
