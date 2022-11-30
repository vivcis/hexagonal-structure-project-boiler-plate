package server

import (
	"log"
	"vivicis/github.com/notification-service/internal/adapters/repository"
	"vivicis/github.com/notification-service/internal/core/helpers"
)

func Run() (*repository.Db, error) {
	err := helpers.Load()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	db, err := repository.NewDatabaseFactory(repository.InstanceMongo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
