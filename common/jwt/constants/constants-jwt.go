package constants

import "github.com/merq-rodriguez/twitter-clone-backend-go/common/config"

var viper, _ = config.Settings()

/*SECRET_KEY JWT */
var SECRET_KEY = viper.GetString("jwt.secret-key")

/*EXPIRES_IN JWT */
var EXPIRES_IN = viper.GetInt64("jwt.expiresin")
