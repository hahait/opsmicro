package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"ops.was.ink/opsmicro/auth/basic/config"
	"time"
)

// claims 中的信息; 可以基于实际需要添加更多的信息, 比如用户的角色等
type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JwtAuth struct {}

// 生成 token
func (j *JwtAuth) GenerateJwtToken(username string) (token string, err error){
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.TokenExpiredTime).Unix(),
		},
	})
	token, err = t.SignedString([]byte(config.TokenSecret))
	return
}

// 验证 token
func (j *JwtAuth) ValidateJwtToken(tokenstr string) (bool, error){
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenSecret), nil
	})

	if !token.Valid {
		return false, fmt.Errorf("token 验证失败, 错误信息: %s , 请重新登陆...", err.Error())
	}

	return true, nil

}