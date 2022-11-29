package repository

import "vivicis/github.com/notification-service/internal/ports"

type AWS struct {
}

func NewAWS() ports.AWSRepository {
	return &AWS{}
}
