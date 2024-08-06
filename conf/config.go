package conf

import "github.com/spf13/viper"

type config struct {
	App    *app      `mapstructure:"app"`
	Db     *database `mapstructure:"database"`
	Redis  *redis    `mapstructure:"redis"`
	Logger *logger   `mapstructure:"logger"`
	File   *file     `mapstructure:"file"`
}

var Cfg config

func InitConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(err)
	}
}
