package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
	"math/rand"
)

func GifHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rand.Seed(time.Now().Unix())
		files, err := ioutil.ReadDir("files/cats/gifs")
		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed to read gifs",
			})
			panic(err)
		}
		c.JSON(200, gin.H{
			"gif": files[rand.Intn(len(files))],
		})
	}
}