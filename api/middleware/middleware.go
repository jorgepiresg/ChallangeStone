package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const secrectKey = "my_secrect_key"

var config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(jwtCustomClaims)
	},
	SigningKey: []byte(secrectKey),
}

type jwtCustomClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func CreateJwtToken(ID string) (string, error) {

	claims := &jwtCustomClaims{
		ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secrectKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

var Private = echojwt.WithConfig(config)

func GetIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	return claims.ID
}
