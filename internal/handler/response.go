package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Logger.Error().Msg(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
