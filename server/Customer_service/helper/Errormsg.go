package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowError(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": msg,
	})
}
