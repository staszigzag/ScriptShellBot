package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	Start  string
	Finish string
	Help   string
}

type Errors struct {
	Default        string
	UnknownCommand string
}

type Config struct {
	TelegramToken string
	Scripts       map[string]string
	Messages      Messages
	SudoChatId    int64
	Debug         bool
}

// Init populates Config struct with values from config file
func Init(path string) (*Config, error) {
	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseConfigFile(filePath string) error {
	sep := "/"
	path := strings.Split(filePath, sep)
	folder := strings.Join(path[:len(path)-1], sep)
	fileName := path[len(path)-1]
	viper.AddConfigPath(folder)   // folder
	viper.SetConfigName(fileName) // config file name

	return viper.ReadInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.error", &cfg.Messages.Errors); err != nil {
		return err
	}

	return nil
}
