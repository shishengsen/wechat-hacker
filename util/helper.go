package util

import (
	"bytes"
	"encoding/gob"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"time"
)

type Token struct {
	Token      string `json:"token"`
	ExpireTime int64  `json:"expireTime"`
}

// 生成token
func MakeToken(role public.RoleType) (tokenRes *Token, err error) {
	// 获取token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// 有效期七天
	expTime := time.Now().Add(7 * 24 * time.Hour * time.Duration(1)).Unix()
	claims["exp"] = expTime
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(getSecretKey(role)))
	tokenRs := &Token{}
	if err != nil {
		return tokenRs, err
	}
	tokenRs.ExpireTime = expTime
	tokenRs.Token = tokenString
	return tokenRs, nil
}

func ValidateToken(ctx *gin.Context, role public.RoleType) bool {
	token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(getSecretKey(role)), nil
		})
	if err != nil {
		return false
	}
	return token.Valid
}

func ValidateWsToken(req *http.Request, role public.RoleType) bool {
	token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(getSecretKey(role)), nil
		})
	if err != nil {
		return false
	}
	return token.Valid
}

func getSecretKey(st public.RoleType) (secret string) {
	secretKey := ""
	switch st {
	case public.RoleManager:
		secretKey = lib.GetStringConf("base.secrets.manager")
	case public.RoleCustomer:
		secretKey = lib.GetStringConf("base.secrets.customer")
	case public.RoleClient:
		secretKey = lib.GetStringConf("base.secrets.client")
	default:
		secretKey = ""
	}
	return secretKey
}

func CallMethod(method interface{}, a interface{}) {
	// here method is a interface which is a type of func
	fv := reflect.ValueOf(method)
	args := []reflect.Value{reflect.ValueOf(a)}
	fv.Call(args)
}


func EncodeCmd(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeCmd(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}