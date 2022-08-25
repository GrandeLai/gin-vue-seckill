package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserToken struct {
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

var FrontUserExpireDuration = time.Hour
var FrontUserSecretKey = []byte("front_user_token")

var AdminUserExpireDuration = time.Hour * 2
var AdminUserSecretKey = []byte("admin_user_token")

//生成token
func GenToken(userName string, expireDuration time.Duration, secretKey []byte) (string, error) {
	// 创建一个我们自己的声明
	user := UserToken{
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(), // 过期时间
			Issuer:    "gin_vue_seckill",                     // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(secretKey)
}

//认证token
func AuthToken(tokenString string, secretKey []byte) (*UserToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserToken{}, func(token *jwt.Token) (key interface{}, err error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, is_ok := token.Claims.(*UserToken)
	if is_ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token valid err")
}
