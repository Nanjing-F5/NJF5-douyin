package common

import (
	"douyin/global"
	"github.com/golang-jwt/jwt"
	"time"
)

var SigningKey = []byte(global.Config.Jwt.SigningKey)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToke 生成token
func GenerateToke(username string) (string, error) {

	// 创建一个用户的声明
	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60, // 过期时间
			Issuer:    username,                  // 签发人
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的key签名并获得完整的编码后的字符串token
	return token.SignedString(SigningKey)
}

// VerifyToken 验证token
func VerifyToken(tokenString string) error {
	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return err
}
