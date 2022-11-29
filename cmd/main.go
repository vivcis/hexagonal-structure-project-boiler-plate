package main

import (
	"log"
	"vivicis/github.com/notification-service/cmd/server"
)

func main() {
	db, err := server.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Injection(db)
}
