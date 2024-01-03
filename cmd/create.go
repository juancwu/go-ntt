package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func create() *cobra.Command {
	createCmd := &cobra.Command{
		Use:     "create",
		Short:   "Create migration files",
		Long:    "Creates a pair of migration files for up/down migrations. You can select to use sequential (max of 6 digits) or timestamp format. Defaults to timestamp.",
		Example: "ntt create NAME",
		Aliases: []string{"c", "new"},
		Args:    cobra.ExactArgs(1),
		RunE:    handle,
	}

	return createCmd
}

func handle(cmd *cobra.Command, args []string) error {
	_, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}
	return nil
}
