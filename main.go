package main

import (
	"log"
	"usership/config"
	"usership/db"
	"usership/routers"
)

func Init() {
	//load config about database and web server
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	//initialize database
	db.Init()
}
func main() {
	Init()
	s := routers.Service{}
	s.InitAndRun()
}
