package auth

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.io/iv587/goredis-admin/db"
	"time"
)

var (
	secret     = "goredis"
	expireTime = 3600 * 7
)

type jwtClaims struct {
	jwt.StandardClaims
	Id       int
	Name     string
	Password string
}

func Gen(user db.User) (string, error) {
	claims := new(jwtClaims)
	claims.Id = user.Id
	claims.Name = user.Name
	claims.Password = user.Password
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(expireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("网络错误，请重试")
	}
	return signedToken, nil
}

func Verify(strToken string) (bool, *db.User, error) {
	if strToken == "" {
		return false, nil, nil
	}
	token, err := jwt.ParseWithClaims(strToken, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false, nil, errors.New("网络错误，请重试")
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return false, nil, nil
	}
	if err := token.Claims.Valid(); err != nil {
		return false, nil, nil
	}
	return true, &db.User{
		Id:       claims.Id,
		Password: claims.Password,
		Name:     claims.Name,
	}, nil
}
