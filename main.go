package main

import (
	"github.com/JincorTech/backend-cert2addr/app"
	"github.com/JincorTech/backend-cert2addr/config"
)

func main() {
	a := app.Application{}

	a.Initialize(config.GetHttpAuthUsername(), config.GetHttpAuthPassword())

	a.Run(config.GetListenAddress())
}
