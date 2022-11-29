package service

import "vivicis/github.com/notification-service/internal/ports"

type mailerService struct {
	mailerRepository ports.MailerRepository
}

type userService struct {
	userRepository ports.UserRepository
}

type AWSService struct {
	AWSRepository ports.AWSRepository
}

func NewAWSServices(AWSRepository ports.AWSRepository) ports.AWSService {
	return &AWSService{
		AWSRepository: AWSRepository,
	}
}

func NewMailerService(mailerRepository ports.MailerRepository) ports.MailerService {
	return &mailerService{
		mailerRepository: mailerRepository,
	}
}

func NewUserService(userRepository ports.UserRepository) ports.UserService {
	return &userService{
		userRepository: userRepository,
	}
}
