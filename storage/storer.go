package storage

type RedisInterface interface {
	StoreLongUrl2Redis(url string) string
	GetShortUrlFromRedis(key uint64) string
}
