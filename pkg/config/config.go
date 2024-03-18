package config

import (
	"time"

	"github.com/spf13/viper"
)

const ConfigName = "config"
const ConfigType = "yaml"

type Config struct {
	Cars      CarsConfig      `mapstructure:"cars"`
	Stations  StationsConfig  `mapstructure:"stations"`
	Registers RegistersConfig `mapstructure:"registers"`
}

type CarsConfig struct {
	Count          int           `mapstructure:"count"`
	ArrivalTimeMin time.Duration `mapstructure:"arrival_time_min"`
	ArrivalTimeMax time.Duration `mapstructure:"arrival_time_max"`
}

type StationsConfig struct {
	Gas      StationConfig `mapstructure:"gas"`
	Diesel   StationConfig `mapstructure:"diesel"`
	Lpg      StationConfig `mapstructure:"lpg"`
	Electric StationConfig `mapstructure:"electric"`
}

type StationConfig struct {
	Count        int           `mapstructure:"count"`
	ServeTimeMin time.Duration `mapstructure:"serve_time_min"`
	ServeTimeMax time.Duration `mapstructure:"serve_time_max"`
}

type RegistersConfig struct {
	Count         int           `mapstructure:"count"`
	HandleTimeMin time.Duration `mapstructure:"handle_time_min"`
	HandleTimeMax time.Duration `mapstructure:"handle_time_max"`
}

func init() {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
}

func NewConfig() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	c := new(Config)
	if err := viper.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}
