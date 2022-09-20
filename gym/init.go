package gym

import "github.com/spf13/viper"

const (
	ConfigKeyUsername = "username"
	ConfigKeyPassword = "password"
	ConfigKeyDebug    = "debug"
)

func init() {
	viper.SetEnvPrefix("GYM")
	viper.AutomaticEnv()

	viper.SetDefault(ConfigKeyDebug, false)
}
