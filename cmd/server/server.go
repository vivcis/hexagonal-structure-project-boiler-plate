package server

import (
	"log"
	"vivicis/github.com/notification-service/internal/adapters/repository"
	"vivicis/github.com/notification-service/internal/core/helpers"
)

//func redisConnect(url string, password string) *redis.Client {
//
//	logrus.WithField("connection", url).Info("Connecting to Redis DB")
//	client := redis.NewClient(&redis.Options{
//		Addr:     url,
//		Password: password, // no password set
//		DB:       0,        // use default DB
//	})
//	err := client.Ping().Err()
//	if err != nil {
//		logrus.Fatal(err)
//	}
//	return client
//
//}

func Run() (*repository.Db, error) {
	err := helpers.Load()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db, err := repository.NewDatabaseFactory(repository.InstanceMongo)
	//db, err := repository.ConnectDb(&helpers.Config{
	//	DBUser:     helpers.Instance.DBUser,
	//	DBPass:     helpers.Instance.DBPass,
	//	DBHost:     helpers.Instance.DBHost,
	//	DBName:     helpers.Instance.DBName,
	//	DBPort:     helpers.Instance.DBPort,
	//	DBTimeZone: helpers.Instance.DBTimeZone,
	//	DBMode:     helpers.Instance.DBMode,
	//})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//err = repository.MigrateAll(db)
	//if err != nil {
	//	log.Fatal(err)
	//	return nil, err
	//}
	//user := models.User{
	//	Name:  "vivicis",
	//	Email: "",
	//}
	//if err = user.HashPassword(); err != nil {
	//	return nil, err
	//}
	//db.Create(&user)

	return db, nil
}
