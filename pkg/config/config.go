package config

import (
	"errors"

	"github.com/spf13/viper"
)

var errEmptyToken = errors.New("empty token. Please, export TG_TOKEN and OPEN_AI to env")

type Messages struct {
	Responses
	Errors
}

type Config struct {
	TelegramToken  string
	OpenAiApiKey   string
	PollingTimeout int64  `mapstructure:"polling_timeout"`
	DataBaseFile   string `mapstructure:"db_file"`
	PromptStart    string `mapstructure:"prompt_start"`
	Messages       Messages
}

type Responses struct {
	Start        string `mapstructure:"start"`
	LoadingBrief string `mapstructure:"loading_brief"`
}

type Errors struct {
	Default string `mapstructure:"default"`
}

func LoadConfig() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := fromEnv(&cfg); err != nil {
		return nil, err
	}

	if cfg.TelegramToken == "" || cfg.OpenAiApiKey == "" {
		return nil, errEmptyToken
	}
	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
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

func fromEnv(cfg *Config) error {
	if err := viper.BindEnv("tg_token"); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("tg_token")

	if err := viper.BindEnv("openai_key"); err != nil {
		return err
	}
	cfg.OpenAiApiKey = viper.GetString("openai_key")
	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
