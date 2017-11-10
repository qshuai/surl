package controllers

import (
	"surl/storage"

	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

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
