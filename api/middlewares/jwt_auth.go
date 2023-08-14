package middlewares

import (
	"log"
	"net/http"

	jwt_service "github.com/freddiemo/logistics-api/api/security/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeJWT(envs map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) > 0 {
			tokenString := authHeader[len(BEARER_SCHEMA):]
			token, err := jwt_service.NewJWTService(envs).ValidateToken(tokenString)

			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				log.Println("Claims[Id]: ", claims["id"])
				log.Println("Claims[Issuer]: ", claims["iss"])
				log.Println("Claims[IssuedAt]: ", claims["iat"])
				log.Println("Claims[ExpiresAt]: ", claims["exp"])

			} else {
				log.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}

		} else {
			log.Println("Not Authorization header")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
