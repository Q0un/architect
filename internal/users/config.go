package users

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HttpAddress    string `yaml:"http_address"`
	LogFile        string `yaml:"log_file"`
	DbUser         string `yaml:"db_user"`
	DbPassword     string `yaml:"db_password"`
	DbName         string `yaml:"db_name"`
	DbHost         string `yaml:"db_host"`
	DbPort         string `yaml:"db_port"`
	JwtPrivateFile string `yaml:"jwt_private_file"`
	JwtPublicFile  string `yaml:"jwt_public_file"`
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
