package middleware

import (
	"net/http"

	"github.com/IRSHIT033/E-comm-GO-/server/Order_service/grpc_auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		is_authorized := grpc_auth.CheckTokenValidityViaGRPC(authHeader)

		if is_authorized {
			c.Next()
			return
		} else {
			c.JSON(http.StatusUnauthorized, "token is not valid")
			c.Abort()
			return
		}

	}

}
