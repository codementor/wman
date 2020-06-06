package weather

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// GetConfig retrieves a WeatherConfig from a file
func GetConfig(file string) (*WeatherConfig, error) {

	wc := &WeatherConfig{}
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, wc)
	if err != nil {
		return nil, err
	}

	return wc, nil
}
