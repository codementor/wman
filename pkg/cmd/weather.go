package cmd

import (
	"github.com/spf13/cobra"
)

const weatherDesc = `
This command consists of multiple sub-commands to interact with weather for open weather map.

There is a option for retrieving weather for a single city or for a group of cities.
`

const (
	weatherGetExample = `  # Retrieving weather for a city
  wman weather get [city]
`
)

type weatherOptions struct {
	config string
}

// newWeatherCmd returns a new initialized instance of the weather sub command
func newWeatherCmd() *cobra.Command {
	opts := &weatherOptions{}
	cmd := &cobra.Command{
		Use:   "weather",
		Short: "Displays different options for weather",
		Long:  weatherDesc,
	}
	cmd.PersistentFlags().StringVar(&opts.config, "config", "config.yaml", "The config to use for weather.")
	cmd.AddCommand(newWeatherGetCmd())
	return cmd
}

// newWeatherGetCmd creates a command that shows the weather for a city.
func newWeatherGetCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "get",
		Short:   "Displays the weather for a city.",
		Example: weatherGetExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
