package storage

import (
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var currentIndex uint64 = uint64(viper.GetInt("start"))

// StoreLongUrl2Redis store a long url to redis server
func (r *RedisStorage) StoreLongUrl2Redis(url string) string {
	p := r.pool.Get()
	defer p.Close()

	index := atomic.AddUint64(&currentIndex, 1)
	atomic.StoreUint64(&currentIndex, index)
	_, err := p.Do("SET", index, url)
	if err != nil {
		fmt.Println("\033[;31mstore long url error:", err, "\033[;0m")
		return ""
	}
	return strconv.FormatUint(index, 36)
}

// GetShortUrlFromRedis acquire a long url with the short one
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
