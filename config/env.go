package config

import (
	"log"
	"strconv"
)

type Env string

func (e Env) IsProd() bool {
	return e == "production"
}

func (e Env) IsDev() bool {
	return e == "development"
}

func (e Env) IsLocal() bool {
	return e == "local"
}

func (e Env) String() string {
	return string(e)
}

func (e Env) Int() int {
	conf, err := strconv.Atoi(e.String())
	if err != nil {
		log.Println(err)
	}

	return conf
}

func (e Env) Int64() int64 {
	conf, err := strconv.ParseInt(e.String(), 10, 64)
	if err != nil {
		log.Println(err)
	}

	return conf
}
