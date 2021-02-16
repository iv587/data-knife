package auth

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secret     = "goredis"
	expireTime = 3600 * 7
)

type jwtClaims struct {
	jwt.StandardClaims
	Id int
}

func Gen(id int) (string, error) {
	claims := new(jwtClaims)
	claims.Id = id
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(expireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("网络错误，请重试")
	}
	return signedToken, nil
}

func Verify(strToken string) (bool, int, error) {
	if strToken == "" {
		return false, 0, nil
	}
	token, err := jwt.ParseWithClaims(strToken, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false, 0, errors.New("网络错误，请重试")
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return false, 0, nil
	}
	if err := token.Claims.Valid(); err != nil {
		return false, 0, nil
	}
	return true, claims.Id, nil
}
