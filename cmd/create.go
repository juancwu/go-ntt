package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/juancwu/go-ntt/util"
	"github.com/spf13/cobra"
)

func create() *cobra.Command {
	createCmd := &cobra.Command{
		Use:     "create",
		Short:   "Create migration files",
		Long:    "Creates a pair of migration files for up/down migrations. You can select to use sequential (max of 6 digits) or timestamp format. Defaults to timestamp.",
		Example: "ntt create NAME",
		Aliases: []string{"new"},
		Args:    cobra.ExactArgs(1),
		RunE:    handle,
	}

	return createCmd
}

func handle(cmd *cobra.Command, args []string) error {
	tz, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	timestamp := time.Now().In(tz)

	name := args[0]

	version := strconv.FormatInt(timestamp.Unix(), 10)

	if err = os.MkdirAll(config.Source, os.ModePerm); err != nil {
		return err
	}

	for _, dir := range [2]string{"up", "down"} {
		filename := filepath.Join(config.Source, fmt.Sprintf("%s_%s_%s%s", version, name, dir, config.Ext))
		util.Log().Debug("New migration file", "filename", filename)

		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			return err
		}

		err = file.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
