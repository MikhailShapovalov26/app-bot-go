package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Telegram struct {
		Token    string `yaml:"bot_token", envconfig:"BOT_TOKEN"`
		Chat_id  string `yaml:"chat_id", envconfig:"CHAT_ID"`
		User     string `yaml:"user", envconfig:"USER"`
		Password string `yaml:"password", envconfig:"PASSWORD"`
		Port     string `yaml:"port", envconfig:"SERVER_PORT"`
		Host     string `yaml:"host", envconfig:"SERVER_HOST"`
	} `yaml:"telegram"`
}

func LoadFile(path string) (*Config, error) {

	var cfg Config

	readFile(&cfg, path)
	readEnv(&cfg)

	return &cfg, nil

}

func readFile(cfg *Config, path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decode := yaml.NewDecoder(f)
	err = decode.Decode(cfg)
	if err != nil {
		fmt.Println(err)
	}
}
func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		fmt.Println(err)
	}

}
