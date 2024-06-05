package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 4000000000
// time.Duration(time.Second*60*6)
func main() {
	strToken, err := NewUserToken("123456", "xx", time.Duration(500000000))
	if err != nil {
		fmt.Errorf("new user token error,err:%s", err.Error())
		return
	}
	fmt.Printf("token:%v\n", strToken)

	//check
	err = CheckHttpToken(strToken)
	if err != nil {
		fmt.Errorf("check user auth internal parse with claims error,err:%s\n", err.Error())
		return
	}
	fmt.Printf("check token ok\n")
}

const (
	TOKEN_EFFECTIVE_RANDOM_SECRET_FACTOR = "qkmsverysecret"
)

type Claims struct {
	jwt.StandardClaims

	SdkId     string `json:"sdk_id"`
	SubSdkId  string `json:"sub_sdk_id"`
	AppId     string `json:"app_id"`
	UserId    string `json:"user_id"`
	SessionId string `json:"session_id"`
}

func CheckHttpToken(tokenStr string) error {

	var clientClaims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			fmt.Errorf("ParseWithClaims err")
			return nil, fmt.Errorf("ParseWithClaims err")
		}
		fmt.Printf("xxxxxxxxx\n")
		return []byte(TOKEN_EFFECTIVE_RANDOM_SECRET_FACTOR), nil
	})
	fmt.Printf("ooooooooooooooo\n")
	if err != nil {
		fmt.Printf("check user auth internal parse with claims error,err:%s\n", err.Error())
		return err
	}
	fmt.Printf("wwwwwwwwwwwwwwwwwwwww\n")
	if !token.Valid {
		fmt.Printf("token.NotValid\n")

		return fmt.Errorf("invalid token")
	}

	//if clientClaims.ExpiresAt{}

	fmt.Printf("clientClaims.SessionId：%v\n", clientClaims.SessionId)
	fmt.Printf("clientClaims.ExpAt：%v\n", clientClaims.ExpiresAt)
	fmt.Printf("time Now：%v\n", time.Now().Unix())
	return nil
}

func NewUserToken(sessionId string, tokenType string, expTime time.Duration) (strToken string, err error) {
	fmt.Printf("newUserToken called with params sessionId:%+v", sessionId)
	nowTime := time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":        "zdlz",                         //issuer:签发方
		"nbf":        nowTime,                        //Not (valid) Before:生效时间
		"exp":        time.Now().Add(expTime).Unix(), //Expiration Time:过期时间
		"user_id":    sessionId,
		"sdk_id":     "",
		"type":       tokenType,
		"session_id": sessionId,
	})
	strToken, err = token.SignedString([]byte(TOKEN_EFFECTIVE_RANDOM_SECRET_FACTOR))
	if err != nil {
		fmt.Errorf("new user token error,err:%s,sessionId:%s\n", err.Error(), sessionId)
		return "", err
	}

	fmt.Printf("newUserToken called with params sessionId:%+v,ret token:%s\n", sessionId, strToken)
	return
}
