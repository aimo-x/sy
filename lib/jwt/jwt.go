package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = "asfdeisaufhwe"

// WeChat jwt
type WeChat struct {
}

// Token 生成 Token 传入的 uid, ouid 请加密 iss
func (w *WeChat) Token(unionid, openid, appid, nickname string) (tokenString string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"union_id":  unionid,
		"open_id":   openid,
		"app_id":    appid,
		"nick_name": nickname,
		"nbf":       time.Now().Unix(),
		"exp":       time.Now().Add(3600 * time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString([]byte(secret))
	return
}

// Verify 验证 token 有效性 (加密后的user_id)
func (w *WeChat) Verify(tokenString string) (unionid, openid, appid, nickname string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		unionid, openid, appid, nickname = claims["union_id"].(string), claims["open_id"].(string), claims["app_id"].(string), claims["nick_name"].(string)
		return
	}
	err = errors.New("JWT ERROR")
	return
}
