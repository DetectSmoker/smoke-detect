package main

import (
	"log"
	"smoke-detect/config"
	"smoke-detect/internal/routers"
)

func main() {
	cf := config.InitConfig()
	log.Println(cf.BucketName)

	router := routers.InitRouter(cf)
	router.Run()
}