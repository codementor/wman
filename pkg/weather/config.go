package weather

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// GetConfig retrieves a Config from a file
func GetConfig(file string) (*Config, error) {

	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist")
	}
	if err != nil {
		return nil, err
	}
	wc := &Config{}
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
