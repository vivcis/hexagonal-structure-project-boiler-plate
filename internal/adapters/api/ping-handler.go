package api

import (
	"github.com/gin-gonic/gin"
	"vivicis/github.com/notification-service/internal/core/helpers"
)

func (u *HTTPHandler) PingHandler(c *gin.Context) {
	// healthcheck
	helpers.JSON(c, "pong", 200, nil, nil)
}
