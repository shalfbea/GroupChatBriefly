package config

import (
	"github.com/spf13/viper"
)

const ()

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
