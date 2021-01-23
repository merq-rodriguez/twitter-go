package config

import (
	"fmt"

	"github.com/spf13/viper"
)

/*Environment struct type*/
type Environment struct {
	port        int
	uriDatabase string
}

var Configuration Environment

/*Settings function*/
func Settings() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return v, err
	}
	return v, nil
}
