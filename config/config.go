package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/caarlos0/env"
	"gopkg.in/yaml.v2"
)

type Config struct{
	Port string
	Db Db
}

type Db struct {
	Port string `env:"DB_PORT"`
	Name string `env:"DB_NAME"`
	User string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

var config *Config
var once sync.Once

func Get() *Config {
	return config
}

func init() {
	once.Do(func(){
		config = &Config{}
		goenv := os.Getenv("GO_ENV")
		filepath := fmt.Sprintf("./config/%v.yml", goenv)
		buf, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatalln("failed to load config yaml err:", err)
		}

		fmt.Println("config:", config)
		err = yaml.Unmarshal(buf, config)
		if err != nil {
			log.Fatalln("failed to Unmarshal yaml err: ", err)
		}

		if err := env.Parse(&config.Db); err != nil {
			fmt.Printf("failed to load Db error: %s", err)
		}
	})
}