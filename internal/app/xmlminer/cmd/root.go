package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "xmlminer",
		Short: "Mine Burp Suite XML for hidden secrets",
	}
	rootCmd.AddCommand(NewExtractCmd())
	rootCmd.SilenceErrors = true
	return rootCmd
}
