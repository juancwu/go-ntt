package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

const (
	cfgFile    = ".gonttrc"
	cfgType    = "yaml"
	timeFormat = "20060102150405"
	timezone   = "UTC"

	dirFlagUsage = "This is the destination the migrations are located. This should be relative to the CWD or an absolute path"
)

func Execute() error {
	rootCmd := &cobra.Command{
		Use:           "ntt",
		Short:         "A simple SQL migration management tool",
		Long:          "Ntt is a simple SQL migration management tool like any other out there in the wild. Ntt is inspired by how Drizzle ORM handles tables per project by using a prefix, which allows multiple projects under the same database for quick development.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(create())
	rootCmd.AddCommand(migrate())

	return rootCmd.ExecuteContext(context.Background())
}
