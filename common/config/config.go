package config

import (
	"github.com/spf13/viper"
)

/*Settings function*/
func Settings() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return v, err
	}
	return v, nil
}
