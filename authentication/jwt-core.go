package authentication

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type JwtClaim struct {
	Username           string `json:"username"`
	Email              string `json:"email"`
	jwt.StandardClaims        // Info of Token: IssuedAt, ExpiresAt, Issuer
}

func GenerateJwt(email, username string) (tokenStr string, err error) {
	expireTime := time.Now().Add(15 * time.Minute)

	claim := &JwtClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err = token.SignedString(jwtKey)
	return
}

func ValideToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}

	claim, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claim.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
