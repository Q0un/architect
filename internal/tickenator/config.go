package tickenator

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Name         string `yaml:"name"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
}

type Config struct {
	GrpcAddress string   `yaml:"grpc_address"`
	LogFile     string   `yaml:"log_file"`
	Db          DbConfig `yaml:"db"`
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
