package constants

import "github.com/merq-rodriguez/twitter-go/common/config"

func GetURLApi() string {
	viper, _ := config.Settings()
	return viper.GetString("http.url")
}

func GetUploatPathAvatar() string {
	viper, _ := config.Settings()
	return viper.GetString("uploads.avatars")
}

func GetUploatPathBanner() string {
	viper, _ := config.Settings()
	return viper.GetString("uploads.banners")
}
