package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "neo-cli",
		Short: "A youtube downloader",
		Long:  `Neo is a youtube downloader that can download videos and save them in a folder of your choice and also uploads the metadata to a Postgres database.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
