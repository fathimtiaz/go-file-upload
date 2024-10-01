package main

import (
	"flag"
	"go-file-upload/app/api"
)

func main() {
	app := flag.String("app", "api", "run api")

	flag.Parse()

	if *app == "api" {
		api.Run()
	} else if *app == "cron" {

	}
}
