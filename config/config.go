package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
}

func InitConfig() (dsn string, err error) {
	var config struct {
		DB DBConfig `json:"db"`
	}

	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal("error opening 'config.json' file")
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("error reading file contents")
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("error unmarshaling data to config")
	}

	dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DBName)

	return dsn, nil
}
