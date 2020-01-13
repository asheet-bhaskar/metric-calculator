package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  ``,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
