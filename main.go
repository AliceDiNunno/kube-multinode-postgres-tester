package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	LoadEnv()

	gormConfig := LoadGormConfiguration()

	db := StartGormDatabase(gormConfig)

	if db == nil {
		panic("Database connection failed")
	}

	hostname, err := os.Hostname()

	var msg string

	if err != nil {
		msg = "Could not get hostname"
	} else {
		msg = hostname
	}

	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hostname": msg,
		})
	})
	r.Run()
	//ginConfig := LoadGinConfiguration()
}
