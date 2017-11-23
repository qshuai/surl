package storage

import (
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var current_index uint64 = uint64(viper.GetInt("start"))

func (r *RedisStorage) StoreLongUrl2Redis(url string) string {
	p := r.pool.Get()
	defer p.Close()

	index := atomic.AddUint64(&current_index, 1)
	atomic.StoreUint64(&current_index, index)
	_, err := p.Do("SET", index, url)
	if err != nil {
		fmt.Println("\033[;31mstore long url error:", err, "\033[;0m")
		return ""
	}
	return strconv.FormatUint(index, 36)
}

func (r *RedisStorage) GetShortUrlFromRedis(key uint64) string {
	p := r.pool.Get()
	defer p.Close()

	lurl, err := redis.String(p.Do("GET", key))
	if err != nil {
		fmt.Println("\033[;31mget long url error:", err, "\033[;0m")
		return ""
	}
	return lurl
}
