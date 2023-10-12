package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(userID string) (string, error) {
	// tokenの作成
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// claimsの設定
	token.Claims = jwt.MapClaims{
		"user": userID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(), // 有効期限を指定
	}

	// 署名
	var secretKey = os.Getenv("SECRET_KEY") // 任意の文字列
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// jwtの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil // CreateTokenにて指定した文字列を使います
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
