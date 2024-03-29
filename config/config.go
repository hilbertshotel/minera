package config

import (
	"os"
	"encoding/json"
	"time"
)

type Templates struct {
	Editor string
	Catalog string
}

type Config struct {
    ConnStr string
	HostAddr string
	Tmp Templates
	MaxAtt int
	ImgDir string
	CookieName string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

// initiate new config struct with default values
func New() *Config {
	return &Config{
		ConnStr: "user=postgres dbname=minera_catalog sslmode=disable host=/run/postgresql",
		HostAddr: "127.0.0.1:5252",
		Tmp: Templates{
			Editor: "templates/editor/*",
			Catalog: "templates/catalog/*",
		},
		MaxAtt: 10,
		ImgDir: "images/",
		CookieName: "minera",
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
}

// parse config file into config struct
func (cfg *Config) Parse() error {
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg);err != nil {
		return err
	}

	return nil
}