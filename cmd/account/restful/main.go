package main

import (
	"flag"
	"log"
)

var path = flag.String("c", "./configs/account/restful/local.yaml", "set config file path")

func init() {
	flag.Parse()
}

// @title IRent API
// @version 0.0.1
// @description API for IRent
//
// @contact.name sean.zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @BasePath /api
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	svc, err := CreateService(*path)
	if err != nil {
		log.Fatal(err)
	}

	err = svc.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = svc.AwaitSignal()
	if err != nil {
		log.Fatal(err)
	}
}
