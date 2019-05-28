package experiment

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func ReadConfig() {
	initializeViper()

	viper.SetDefault("hostname", "0.0.0.0")
	viper.SetDefault("port", 4000)
	viper.Set("host", fmt.Sprintf("%s:%d", viper.Get("hostname"), viper.Get("port")))

	viper.SetDefault("database.name", "playground")
	viper.SetDefault("database.host", "127.0.0.1")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.username", "root")
	viper.SetDefault("database.connection_string", buildConnectionString())
}

func initializeViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Unable to read config: %s\n", err))
	}
}

func buildConnectionString() string {
	host := viper.Get("database.host")
	name := viper.Get("database.name")
	password := viper.Get("database.password")
	username := viper.Get("database.username")

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", username, password, host, name)
}
