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
func (w *WeChat) Token(openid, appid, nickname string) (tokenString string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"openid":   openid,
		"appid":    appid,
		"nickname": nickname,
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(36000 * time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString([]byte(secret))
	return
}

// Parse ...
func (w *WeChat) Parse(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(secret), nil
}

// Verify 验证 token 有效性 (加密后的user_id)
func (w *WeChat) Verify(tokenString string) (openid, appid, nickname string, err error) {
	token, err := jwt.Parse(tokenString, w.Parse)
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// unionid, openid, appid, nickname = claims["union_id"].(string), claims["open_id"].(string), claims["app_id"].(string), claims["nick_name"].(string)
		return claims["openid"].(string), claims["appid"].(string), claims["nickname"].(string), nil
	}
	err = errors.New("JWT ERROR")
	return
}

// Wxmp jwt
type Wxmp struct {
}

// Token 生成 Token 传入的 uid, ouid 请加密 iss
func (w *Wxmp) Token(openid, userid, campaignid string) (tokenString string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"oid": openid,
		"uid": userid,
		"cid": campaignid,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(36000 * time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString([]byte(secret))
	return
}

// Parse ...
func (w *Wxmp) Parse(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	return []byte(secret), nil
}

// Verify 验证 token 有效性 (加密后的user_id)
func (w *Wxmp) Verify(tokenString string) (openid, userid, campaignid string, err error) {
	token, err := jwt.Parse(tokenString, w.Parse)
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// unionid, openid, appid, nickname = claims["union_id"].(string), claims["open_id"].(string), claims["app_id"].(string), claims["nick_name"].(string)
		return claims["oid"].(string), claims["uid"].(string), claims["cid"].(string), nil
	}
	err = errors.New("JWT ERROR")
	return
}
