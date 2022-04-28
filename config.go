package main

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

// Gorm Config

type GormConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func LoadGormConfiguration() GormConfig {
	return GormConfig{
		Host:     RequireEnvString("DB_HOST"),
		Port:     RequireEnvInt("DB_PORT"),
		User:     RequireEnvString("DB_USER"),
		Password: RequireEnvString("DB_PASSWORD"),
		DbName:   RequireEnvString("DB_NAME"),
	}
}

// Gin Config

type GinConfig struct {
	ListenAddress string
	Port          int
	TlsEnabled    bool
	Prefix        string
}

func LoadGinConfiguration() GinConfig {
	prefix, err := GetEnvString("API_PREFIX")

	if err != nil {
		prefix = "/"
	}

	return GinConfig{
		ListenAddress: RequireEnvString("LISTEN_ADDRESS"),
		Port:          RequireEnvInt("PORT"),
		TlsEnabled:    RequireEnvBool("TLS_ENABLED"),
		Prefix:        prefix,
	}
}

// Main Config
func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found.")
	}
}

func GetEnvString(name string) (string, error) {
	envVariable, exists := os.LookupEnv(name)

	if exists == false {
		log.Println("Env variable: ", name, " is not found")
		return "", errors.New("env variable not found")
	}

	if strings.Contains(name, "SECRET") || strings.Contains(name, "KEY") || strings.Contains(name, "TOKEN") || strings.Contains(name, "PASSWORD") {
		log.Println("[", name, "] = ****")
	} else {
		log.Println("[", name, "] = ", envVariable)
	}

	return envVariable, nil
}

func GetEnvStringOrDefault(name string, defaultValue string) string {
	envVariable, err := GetEnvString(name)

	if err != nil {
		return defaultValue
	}

	return envVariable
}

func GetEnvInt(name string) (int, error) {
	value, err := GetEnvString(name)

	if err != nil {
		return 0, err
	}

	envVariable, err := strconv.Atoi(value)

	if err != nil {
		log.Println(name, " value is invalid, should be integer")
		return 0, errors.New("value is not int")
	}

	return envVariable, nil
}

func GetEnvIntOrDefault(name string, defaultValue int) int {
	envVariable, err := GetEnvInt(name)

	if err != nil {
		return defaultValue
	}

	return envVariable
}

func GetEnvBool(name string) (bool, error) {
	value, err := GetEnvString(name)

	if err != nil {
		return false, err
	}

	return value == "true" || value == "TRUE" || value == "1" || value == "yes" || value == "YES", nil
}

func GetEnvBoolOrDefault(name string, defaultValue bool) bool {
	envVariable, err := GetEnvBool(name)

	if err != nil {
		return defaultValue
	}

	return envVariable
}

func RequireEnvString(name string) string {
	value, err := GetEnvString(name)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return value
}

func RequireEnvInt(name string) int {
	envVariable, err := GetEnvInt(name)

	if err != nil {
		log.Fatalln(name, " value is invalid")
	}

	return envVariable
}

func RequireEnvBool(name string) bool {
	value, err := GetEnvBool(name)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return value
}
