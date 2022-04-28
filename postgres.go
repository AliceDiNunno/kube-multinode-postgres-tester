package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func StartGormDatabase(config GormConfig) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)
	db, err := gorm.Open(pg.Open(psqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	return db
}

type HostnameCount struct {
	gorm.Model

	Hostname string `gorm:"primary_key"`
	Count    int    `gorm:"default:0"`
}

func CreateDB(db *gorm.DB) {
	db.AutoMigrate(&HostnameCount{})
}

func IncrementHostname(db *gorm.DB, hostname string) {
	hostnameExists := HostnameCount{}
	err := db.Where("hostname = ?", hostname).First(&hostnameExists)

	if err.Error != nil {
		spew.Dump(err.Error.Error())
		db.Create(&HostnameCount{Hostname: hostname, Count: 1})
	} else {
		db.Model(&hostnameExists).Update("count", hostnameExists.Count+1)
	}
}

func ListHostnames(db *gorm.DB) []HostnameCount {
	var hostnames []HostnameCount
	db.Find(&hostnames)

	return hostnames
}
