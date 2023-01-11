package config

import (
	"flag"
	"os"

	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/util"
	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	BonusUrl  string `yaml:"bonus_url"`
	FlightUrl string `yaml:"flight_url"`
	TicketUrl string `yaml:"ticket_url"`
}

type ServerConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type Config struct {
	Service ServiceConfig
	Server  ServerConfig
}

func LoadConfig() (c *Config, err error) {
	filename := flag.String("configFile", "config.yaml", "Config file (default: config.yaml)")

	data, err := os.ReadFile(*filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return
	}

	c.Server.Host = util.GetEnv("HOST", c.Server.Host)
	c.Server.Port = util.GetEnv("PORT", c.Server.Port)

	return
}
