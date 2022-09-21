package gym

import "github.com/spf13/viper"

const (
	ConfigKeyUsername = "username"
	ConfigKeyPassword = "password"
	ConfigKeyDebug    = "debug"
	ConfigKeyStore    = "store"
)

func init() {
	viper.SetEnvPrefix("GYM")
	viper.AutomaticEnv()
	viper.SetDefault(ConfigKeyStore, ".")
	viper.SetDefault(ConfigKeyDebug, false)
}
