package constants

import (
	"strconv"

	"github.com/merq-rodriguez/twitter-go/common/config"
)

var viper, _ = config.Settings()

/*SecretKey JWT */
var SecretKey = viper.GetString("jwt.secret-key")

/*ExpiresIn JWT */
var ExpiresIn = strconv.Itoa(viper.GetInt("jwt.expiresin"))
