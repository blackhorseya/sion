package main

import (
	"flag"
	"log"
)

var path = flag.String("c", "./configs/rental/cronjob/local.yaml", "set config file path")

func init() {
	flag.Parse()
}

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
