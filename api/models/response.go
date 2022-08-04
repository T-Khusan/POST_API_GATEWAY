package models

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ResponseModel ...
type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type StatusRespoce struct {
	Status string `json:"status"`
}

func NewErrorResponce(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorMessage{message})
}
