package config

import (
	"log"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Env  Env `envconfig:"APP_ENV" required:"true"`
		Port Env `envconfig:"APP_PORT" required:"true"`
	}
	DB struct {
		ConnStr Env `envconfig:"DB_CONNSTR" required:"true"`
	}
	FIle struct {
		MaxSize   Env `envconfig:"FILE_MAX_SIZE" required:"true"`
		ChunkSize Env `envconfig:"FILE_CHUNK_SIZE" required:"true"`
	}
}

func LoadDefault() *Config {
	return load(".env")
}

func load(env string) *Config {
	var config Config

	readEnv(env)
	err := envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}
	return &config
}

func readEnv(env string) {
	err := godotenv.Overload(getSourcePath() + "/../" + env)
	if err != nil {
		log.Print(err)
	}
}

func getSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
