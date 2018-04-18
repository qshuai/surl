package storage

// RedisInterface define data interface
type RedisInterface interface {
	StoreLongUrl2Redis(url string) string
	GetShortUrlFromRedis(key uint64) string
}
