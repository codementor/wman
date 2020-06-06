package cmd

import (
	"github.com/spf13/cobra"
)

// NewWmanCmd creates a new root command for wman
func NewWmanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "wman",
		Short:        "Weather Man CLI",
		Long:         `Weather Man (wman) CLI for capturing weather information.`,
		SilenceUsage: true,
		Example: `  # Run print function
  wman print nofluff
`,
	}
	cmd.AddCommand(newPrintCmd())
	cmd.AddCommand(newDogYearCmd())
	cmd.AddCommand(newLintCmd())
	return cmd
}
