package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/yuri7030/final-project/internal/api/common"
	"github.com/yuri7030/final-project/internal/api/config"
	"github.com/yuri7030/final-project/internal/api/handlers"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetValue("JWT_KEY")), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		var authJwt common.AuthJWT
		err = mapstructure.Decode(token.Claims, &authJwt)
		if err != nil {
			panic(err)
		}
		c.Set("user", authJwt)

		c.Next()
	}
}

func BlacklistMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			common.ResponseError(c, http.StatusUnauthorized, "Missing Authorization header", nil)
			c.Abort()
			return
		}

		tokenParts := strings.Split(authorizationHeader, " ")
		token := tokenParts[0]

		if handlers.TokenBlacklist[token] {
			common.ResponseError(c, http.StatusUnauthorized, "Token revoked", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
