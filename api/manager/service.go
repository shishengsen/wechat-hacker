package manager

import (
	"github.com/e421083458/gin_scaffold/api"
	"github.com/e421083458/gin_scaffold/dao"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/e421083458/gin_scaffold/util"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {

}

func (s *Service) CreateCustomer(input *api.RegisterInput, ctx *gin.Context) (t *util.Token, err error){
	// 校验用户是否存在
	userDao := &dao.User{}
	user, err := userDao.FindByPhone(ctx, input.Phone); if err == nil {
		middleware.ResponseError(ctx, 2001, errors.New("用户已存在"))
		return
	}
	// 生成token
	token, err := util.MakeToken(input.Role); if err != nil {
		middleware.ResponseError(ctx, 2001, err)
	}
	// 生成密码
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost); if err != nil{
		middleware.ResponseError(ctx, 2007, err)
	}
	user = &dao.User{
		UserName:input.UserName,
		Role:input.Role,
		Phone:input.Phone,
		Password:string(hashPwd),
		CustomerId:input.CustomerId,
		CompanyIds:input.CompanyIds,
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
		ApiToken:token.Token,
	}
	if err := user.Save(ctx); err != nil {
		middleware.ResponseError(ctx, 2007, err)
	}
	return token, nil
}

