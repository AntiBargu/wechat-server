package http

import (
	"github.com/gin-gonic/gin"
)

type ContentHandler interface {
	Parser(body []byte, c *gin.Context)
}

type ContentHanlderMap interface {
	GetHandler(string) ContentHandler
	SetHandler(string, ContentHandler)
}
