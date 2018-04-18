package controllers

import (
	"surl/storage"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// StoreUrl store a long url with short one
func StoreUrl(c *gin.Context) {
	origin := c.PostForm("long_url")
	fmt.Println(origin)
	str := storage.RedisPool.StoreLongUrl2Redis(origin)
	if str == "" {
		c.JSON(500, gin.H{
			"surl": "",
			"msg":  "failed",
		})
	} else {
		c.JSON(200, gin.H{
			"surl": str,
			"msg":  "success",
		})
	}
}

// GetUrl get the specific origin long url represented by a short one
func GetUrl(c *gin.Context) {
	surl := c.Param("surl")
	fmt.Println(surl)
	key, err := strconv.ParseUint(surl, 36, 0)
	if err != nil {
		fmt.Println("\033[;31murl not validated\033[;0m")
	}
	str := storage.RedisPool.GetShortUrlFromRedis(key)
	fmt.Println(str)
	if str == "" {
		c.JSON(500, gin.H{
			"surl": "",
			"msg":  "failed",
		})
	}
	c.Redirect(302, str)
}
