package app

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("surl: viper component read config file error %s \n", err))
	}
}
