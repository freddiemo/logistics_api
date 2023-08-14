package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	// GenerateToken(id int) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService(envs map[string]string) JWTService {
	return &jwtService{
		secretKey: getSecretKey(envs),
		issuer:    "logistics",
	}
}

func getSecretKey(envs_params map[string]string) string {
	secret := envs_params["JWT_SECRET"]
	if secret == "" {
		secret = "top_secret"
	}
	return secret
}

func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})
}
