package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	HTTPServer httpServer
	Redis      redis
	App        app
}

type app struct {
}

type httpServer struct {
	Host string `envconfig:"HOST" default:"localhost"`
	Port int    `envconfig:"PORT" default:"8443"`
}

type redis struct {
	URI      string `envconfig:"REDIS_URI" default:"localhost:6379"`
	Password string `envconfig:"REDIS_PASSWORD" default:""`
}

var cfg config

func Init() {
	_ = godotenv.Load()
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}
}

func Get() config {
	return cfg
}
