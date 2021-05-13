package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("夏天夏天悄悄过去")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (aToken string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		UserID:   userID,
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串数据
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	//rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	//	ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
	//	Issuer:    "bluebell",
	//}).SignedString(mySecret)

	return

}

// ParseToken 解析JWT
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid token")
	}
	return
}

//func RefreshToken(aToken, rToken string) (newAToken, newERoken string, err error) {
//	if _, err = jwt.Parse(rToken, func(token *jwt.Token) (i interface{}, e error) {
//		return mySecret, nil
//	}); err != nil {
//		return
//	}
//
//	var claims MyClaims
//	_, err = jwt.ParseWithClaims(aToken, &claims, func(token *jwt.Token) (i interface{}, e error) {
//		return mySecret, nil
//	})
//	v, _ := err.(*jwt.ValidationError)
//
//	if v.Errors == jwt.ValidationErrorExpired {
//		return GenToken(claims.UserID, claims.Username)
//	}
//	return
//}
