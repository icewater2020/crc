package config

import (
	"errors"
	"fmt"

	"github.com/code-ready/crc/pkg/crc/config"
	"github.com/spf13/cobra"
)

func configSetCmd(config config.Storage) *cobra.Command {
	return &cobra.Command{
		Use:   "set CONFIG-KEY VALUE",
		Short: "Set a crc configuration property",
		Long:  `Sets a crc configuration property.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("Please provide a configuration property and its value as in 'crc config set KEY VALUE'")
			}
			setMessage, err := config.Set(args[0], args[1])
			if err != nil {
				return err
			}

			if setMessage != "" {
				fmt.Println(setMessage)
			}
			return nil
		},
	}
}
