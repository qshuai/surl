package storage

import (
	"fmt"

	"github.com/spf13/viper"
)

var RedisPool RedisInterface

//获取redis配置信息
func init() {
	redis_config := Redis{
		Host:             viper.GetString("redis.host"),
		Port:             viper.GetInt("redis.port"),
		Password:         viper.GetString("redis.password"),
		DB:               viper.GetInt("redis.db"),
		IdleTimeout:      viper.GetInt("redis.idle_timeout"),
		MaxIdleConns:     viper.GetInt("redis.max_idle_conns"),
		MaxOpenConns:     viper.GetInt("redis.max_open_conns"),
		InitialOpenConns: viper.GetInt("redis.initial_open_conns"),
	}

	var err error
	RedisPool, err = NewRedis(redis_config)
	if err != nil {
		panic(fmt.Errorf("redis: redis init error %s \n", err))
	}

}
