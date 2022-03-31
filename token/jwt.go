package token

import (
	"anya-day/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(uid int) (string, error) {
	lifespan := utils.StringToIntIgnore((utils.GetEnvWithFallback("TOKEN_HOUR_LIFESPAN", "1")))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user_id":    uid,
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(time.Hour * time.Duration(lifespan)).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
