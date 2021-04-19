package conf

import (
	"os"
	"encoding/json"
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
}

// initiate new config struct with default values
func NewConfig() *Config {
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
	}
}

// parse config file into config struct
func (cfg *Config) Parse() error {
	file, err := os.Open("conf/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}