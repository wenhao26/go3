package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTLib struct {
	SetClaims Claims
	Secret    string
}

type Claims struct {
	Data interface{} `json:"data"`
	jwt.StandardClaims
}

// 生成JWT
func (j *JWTLib) GenJWT(data interface{}) string {
	if data != nil {
		j.SetClaims.Data = data
	}
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, j.SetClaims)
	token, err := reqClaims.SignedString([]byte(j.Secret))
	if err != nil {
		return ""
	}

	return token
}

// 校验JWT
func (j *JWTLib) CheckJWT(token string) (*Claims, interface{}) {
	var c Claims

	setToken, err := jwt.ParseWithClaims(token, &c, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.Secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token不正确")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, fmt.Errorf("token已过期")
			} else {
				return nil, fmt.Errorf("token格式错误")
			}
		}
	}

	if setToken != nil {
		if key, ok := setToken.Claims.(*Claims); ok && setToken.Valid {
			return key, nil
		} else {
			return nil, fmt.Errorf("token不正确")
		}
	}

	return nil, fmt.Errorf("token不正确")
}

func main() {
	// JWT
	setClaims := Claims{
		map[string]interface{}{
			"id":   1688,
			"name": "test",
		},
		jwt.StandardClaims{
			Audience:  "接收JWT者",
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Id:        "唯一标识符",
			Issuer:    "发布者",
		},
	}

	jwtLib := JWTLib{
		SetClaims: setClaims,
		Secret:    "12345",
	}
	token := jwtLib.GenJWT(nil)
	fmt.Println("生成的token:", token)

	result, _ := jwtLib.CheckJWT(token)
	fmt.Println("解析的token:", result)
}
