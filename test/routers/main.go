package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	go handler1()
	go handler2()
	time.Sleep(time.Hour)
}
func handler1() {
	r := gin.Default()
	r.GET("/mynginx", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Nginx: Index from 8081",
		})
	})
	r.Run(":8081")
}
func handler2() {
	r := gin.Default()
	r.GET("/mynginx", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Nginx: Index from 8082",
		})
	})
	r.Run(":8082")
}
