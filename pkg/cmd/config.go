package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/codementor/wman/pkg/weather"
)

var (
	configExample = `  # Display current configuration
  wman config`
)

// newConfigCmd returns a new initialized instance of the config sub command
func newConfigCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "config",
		Short:   "Displays the current configuration",
		Example: configExample,
		RunE:    ConfigCmd,
	}

	return cmd
}

// ConfigCmd performs the print config command
func ConfigCmd(cmd *cobra.Command, args []string) error {

	wc, err := weather.GetConfig("config.yaml")
	if err != nil {
		return err
	}
	
	fmt.Println("configuration:")
	fmt.Println("  key:", wc.Key)
	fmt.Println("  unit:", wc.Unit)
	fmt.Println("  cities:")
	for _, city := range wc.Cities {
		fmt.Println("   ", city)
	}
	return nil
}
