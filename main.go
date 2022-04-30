package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func main() {
	println("ENV")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], "=", pair[1])
	}

	LoadEnv()

	gormConfig := LoadGormConfiguration()

	db := StartGormDatabase(gormConfig)

	if db == nil {
		panic("Database connection failed")
	}

	CreateDB(db)

	hostname, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		IncrementHostname(db, hostname)

		c.JSON(200, ListHostnames(db))
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run()
	//ginConfig := LoadGinConfiguration()
}
