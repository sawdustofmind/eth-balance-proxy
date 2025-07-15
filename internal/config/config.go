package config

import (
	"github.com/spf13/viper"

	"github.com/sawdustofmind/eth-balance-proxy/internal/log"
	"github.com/sawdustofmind/eth-balance-proxy/internal/service"
)

type Config struct {
	Logger        *log.Config                  `mapstructure:"logger"`
	Port          int                          `mapstructure:"port"`
	MetricsPort   int                          `mapstructure:"metrics_port"`
	BalanceGetter *service.BalanceGetterConfig `mapstructure:"balance_getter"`
}

func Init() (Config, error) {
	// Load config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
