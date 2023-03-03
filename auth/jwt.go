package auth

import (
	"dk/user"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secret     = "dkdata"
	expireTime = 3600 * 7
)

type jwtClaims struct {
	jwt.StandardClaims
	Id       int
	UserName string
	Password string
}

// Token 生成token
func Token(user user.User) (string, error) {
	claims := new(jwtClaims)
	claims.Id = user.Id
	claims.UserName = user.UserName
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

// Verify 校验token
func Verify(strToken string) (bool, *user.User, error) {
	if strToken == "" {
		return false, nil, nil
	}
	token, err := jwt.ParseWithClaims(strToken, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false, nil, nil
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return false, nil, nil
	}
	if err := token.Claims.Valid(); err != nil {
		return false, nil, nil
	}
	return true, &user.User{
		Id:       claims.Id,
		Password: claims.Password,
		UserName: claims.UserName,
	}, nil
}
