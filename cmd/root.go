package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/juancwu/go-ntt/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cfgFile    = ".gonttrc"
	cfgType    = "yaml"
	timeFormat = "20060102150405"
	timezone   = "UTC"
)

type Config struct {
	Source  string `yaml:"source"`
	Verbose bool   `yaml:"verbose"`
}

var (
	config Config
)

func Execute() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	rootCmd := &cobra.Command{
		Use:   "ntt",
		Short: "A simple SQL migration management tool",
		Long:  "Ntt is a simple SQL migration management tool like any other out there in the wild. Ntt is inspired by how Drizzle ORM handles tables per project by using a prefix, which allows multiple projects under the same database for quick development.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			viper.SetConfigName(cfgFile)
			viper.SetConfigType(cfgType)
			viper.AddConfigPath(cwd)

			err := viper.ReadInConfig()
			if err != nil {
				return fmt.Errorf("config file '%s.%s' not found in [%s]", cfgFile, cfgType, cwd)
			}

			if !viper.IsSet("source") {
				return fmt.Errorf("The 'source' property is missing in the config file")
			}

			if err := viper.Unmarshal(&config); err != nil {
				return err
			}

			util.Log().Debug("Current", "config", config)

			// verify if source is valid or not
			err = util.IsValidPath(config.Source)
			if err != nil {
				return err
			}

			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(create())

	return rootCmd.ExecuteContext(context.Background())
}
