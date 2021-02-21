package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	APIKeys struct {
		Discord string `yaml:"discord,omitempty"`
		Osu     string `yaml:"osu,omitempty"`
	} `yaml:"apiKeys"`
	Commands []string `yaml:"commands"`
}

func init() {
	config := flag.String("c", "", "Specifies path to config file")
	flag.Parse()

	if *config == "" {
		log.Fatal("You need to specify a configuration file")
	}

	err := parseConfig(*config)
	if err != nil {
		log.Fatal(err)
	}
}

func parseConfig(configFile string) error {
	file, err := os.Open(configFile)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.UnmarshalStrict(bytes, &config)
	if err != nil {
		return err
	}

	return nil
}

// GetConfigs returns the global configs of the controller
func GetConfigs() *Config {
	return &config
}
