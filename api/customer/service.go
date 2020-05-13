package customer

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/e421083458/gin_scaffold/dao"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"time"
)

type Service struct {
}

// 用户密码校验
func (s *Service) VerifyPassword(ctx *gin.Context, in *LoginInput) (*dao.User, error) {
	var userDao = &dao.User{}
	user, err := userDao.FindByPhone(ctx, in.Username)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, err
	}
	return user, nil
}

// 生成token
func (s *Service) MakeToken(ctx *gin.Context) (tokenStr string, err error) {
	SecretKey := lib.GetStringConf("app.secret")
	// 获取token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	claims["exp"] = time.Now().Add(7 * 24 * time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}


func (c *Controller) getUserInfoBySess(ctx *gin.Context) (res *LoginResp, err error){
	session := sessions.Default(ctx)
	sessData := session.Get("user")
	resp := &LoginResp{}
	if err := json.Unmarshal([]byte(reflect.ValueOf(sessData).Interface().(string)), resp); err != nil {
		return nil, err
	}
	return resp,nil
}