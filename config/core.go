package config

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	Port            string `json:"port"`
	DefaultLanguage string `json:"default_language"`
	LegacyEndpoit   string `json:"legacy_endpoint"`
	DatabaseType    string `json:"database_type"`
	DatabaseURL     string `json:"json:database_url"`
}

var defaultConfiguration = Configuration{
	Port:            ":8080",
	DefaultLanguage: "english",
}

func (c *Configuration) LoadFromEnv() {

	if lang := os.Getenv("DEFAULT_LANGUAGE"); lang != "" {
		c.DefaultLanguage = lang
	}

	if port := os.Getenv("PORT"); port != "" {
		c.Port = port
	}
}

func (c *Configuration) ParsePort() {

	if c.Port[0] != ':' {
		c.Port = ":" + c.Port
	}

	if _, err := strconv.Atoi(string(c.Port[:1])); err != nil {
		fmt.Printf("invalid port %s", c.Port)
		c.Port = defaultConfiguration.Port
	}
}

func (c *Configuration) LoadFromJSON(path string) error {
	log.Printf("loading configuration from file %s/n", path)

	b, err := os.ReadFile(path)
	if err != nil {
		log.Printf("unable to log file %s/n", err.Error())
		return errors.New("unable to load configuration")
	}

	if err := json.Unmarshal(b, c); err != nil {
		log.Printf("unalbe to parse the file %s/n", err.Error())
		return errors.New("unable to load configuration")
	}

	if c.Port == "" {
		log.Printf("Empty port, reverting to default")
		c.Port = defaultConfiguration.Port
	}

	if c.DefaultLanguage == "" {
		log.Panicf("Empty Language, reverting to default")
		c.DefaultLanguage = defaultConfiguration.DefaultLanguage
	}

	return nil
}

func LoadConfiguration() Configuration {

	cfgfileFlag := flag.String("config_file", "", "load configuration from file")
	portFlag := flag.String("port", "", "set port")

	flag.Parse()
	cfg := defaultConfiguration

	if cfgfileFlag != nil && *cfgfileFlag != "" {
		if err := cfg.LoadFromJSON(*cfgfileFlag); err != nil {
			log.Printf("unable to load configuration form file %s, using default values", *cfgfileFlag)
		}
	}

	cfg.LoadFromEnv()

	if portFlag != nil && *portFlag != "" {
		cfg.Port = *portFlag
	}

	cfg.ParsePort()
	return cfg
}
