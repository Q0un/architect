package users

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type JwtConfig struct {
	PrivateFile string `yaml:"private_file"`
	PublicFile  string `yaml:"public_file"`
}

type Config struct {
	HttpAddress    string    `yaml:"http_address"`
	LogFile        string    `yaml:"log_file"`
	Db             DbConfig  `yaml:"db"`
	Jwt            JwtConfig `yaml:"jwt"`
	TickenatorHost string    `yaml:"tickenator_host"`
}

func LoadConfig(path string) (*Config, error) {
	r, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	conf := &Config{}
	err = yaml.Unmarshal(r, conf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return conf, nil
}
