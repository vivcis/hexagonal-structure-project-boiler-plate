package server

import (
	"vivicis/github.com/notification-service/internal/adapters/api"
	"vivicis/github.com/notification-service/internal/adapters/repository"
	"vivicis/github.com/notification-service/internal/core/helpers"
	"vivicis/github.com/notification-service/internal/core/service"
)

// Injection inject all dependencies
func Injection(db *repository.Db) {
	userRepository := repository.NewUser(db.Mongo)
	userService := service.NewUserService(userRepository)

	mailerRepository := repository.NewMail()
	mailerService := service.NewMailerService(mailerRepository)

	AWSRepository := repository.NewAWS()
	AWSService := service.NewAWSServices(AWSRepository)

	Handler := api.NewHTTPHandler(userService, mailerService, AWSService)
	router := SetupRouter(Handler, userService)

	_ = router.Run(":" + helpers.Instance.Port)
}
