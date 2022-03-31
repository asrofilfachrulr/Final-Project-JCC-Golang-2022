package token

import (
	"anya-day/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// generate jwt token with HS256 alg
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

// validate token
func ValidateToken(c *gin.Context) error {
	token := ExtractToken(c)
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method: %s", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	return nil
}

// extract token from query or authorization header
func ExtractToken(c *gin.Context) string {
	if token := c.Query("token"); token != "" {
		return token
	}

	bearerLine := strings.Split(c.Request.Header.Get("Authorization"), " ")

	if len(bearerLine) == 2 {
		return bearerLine[1]
	}

	return ""
}

// extract user_id from payload of given of jwt token at header or query
func ExtractUID(c *gin.Context) (uint, error) {
	t := ExtractToken(c)
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"].(float64)), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}
