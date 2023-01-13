package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/IRSHIT033/E-comm-GO-/server/User_service/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain_user.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				fmt.Println(userID)
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain_user.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain_user.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
