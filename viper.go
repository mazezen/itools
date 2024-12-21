package itools

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// LoadConfig load config file
func LoadConfig(cf string) {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cf)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	Gm.M = viper.AllSettings()
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		Gm.M = viper.AllSettings()
	})
}

func PrintConfigMap(config map[string]interface{}) {
	for key, value := range config {
		fmt.Printf("%s: %v\n", key, value)
	}
}
