package cmd

import (
	"fmt"
	"sort"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"

	"github.com/codementor/wman/pkg/weather"
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
	cmd.PersistentFlags().StringVarP(&opts.config, "config", "c", "config.yaml", "The config to use for weather.")
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
			file, err := cmd.Flags().GetString("config")
			if err != nil {
				return err
			}
			config, err := weather.GetConfig(file)
			if err != nil {
				return err
			}
			if len(args) < 1 {
				return fmt.Errorf("city must be provided")
			}
			return printCityWeather(config, args)
		},
	}

	return cmd
}

func printCityWeather(config *weather.Config, cities []string) error {
	f, err := weather.New(config)
	if err != nil {
		return err
	}
	mList, err := f.GetCities(cities)
	if err != nil {
		return err
	}

	sort.Sort(mList)

	table := uitable.New()
	table.AddRow("City", "Temp", "Desc")
	for _, m := range mList {
		table.AddRow(m.City, m.Temp, m.Desc)
	}

	fmt.Println(table)
	return nil
}
