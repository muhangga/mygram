package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/muhangga/internal/service"
	"github.com/muhangga/internal/utils"
)

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" {
			resp := utils.JsonError(http.StatusUnauthorized, "Failed to process request", "No Token Found", utils.EmptyObject{})
			c.AbortWithStatusJSON(resp.Status, resp)
			return
		}

		if !strings.Contains(header, "Bearer") {
			resp := utils.JsonError(http.StatusUnauthorized, "Unauthorized", "Unauthorized", utils.EmptyObject{})
			c.AbortWithStatusJSON(resp.Status, resp)
			return
		}

		token := strings.Split(header, " ")[1]
		if token == "" {
			resp := utils.JsonError(http.StatusUnauthorized, "Unauthorized", "Unauthorized", utils.EmptyObject{})
			c.AbortWithStatusJSON(resp.Status, resp)
			return
		}

		validateToken, err := service.NewJwtService().ValidateToken(token)
		if err != nil {
			resp := utils.JsonError(http.StatusUnauthorized, "Token is not valid", "Unauthorized", utils.EmptyObject{})
			c.AbortWithStatusJSON(resp.Status, resp)
			return
		}

		claims := validateToken.Claims.(jwt.MapClaims)

		c.Set("user_id", claims["user_id"])
		c.Set("email", claims["email"])

		c.Next()
	}
}
