package api

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Api struct {
}

func (demo *Api) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	return
}
//
//func (demo *Api) ListPage(c *gin.Context) {
//	listInput := &dto.ListPageInput{}
//	if err := listInput.BindingValidParams(c); err != nil {
//		middleware.ResponseError(c, 2003, err)
//		return
//	}
//	user := &dao.User{}
//	pageInt, err := strconv.ParseInt(listInput.Page, 10, 64)
//	if err != nil {
//		middleware.ResponseError(c, 2004, err)
//		return
//	}
//	if userList, total, err := user.PageList(c, listInput.Name, int(pageInt), 20); err != nil {
//		middleware.ResponseError(c, 2005, err)
//		return
//	} else {
//		m := map[string]interface{}{
//			"list":  userList,
//			"total": total,
//		}
//		middleware.ResponseSuccess(c, m)
//	}
//	return
//}

//func (demo *Api) AddUser(c *gin.Context) {
//	addInput := &dto.AddUserInput{}
//	if err := addInput.BindingValidParams(c); err != nil {
//		middleware.ResponseError(c, 2006, err)
//		return
//	}
//	user := &dao.User{}
//	user.Name = addInput.Name
//	user.Sex = addInput.Sex
//	user.Age = addInput.Age
//	user.Birth = addInput.Birth
//	user.Addr = addInput.Addr
//	user.CreateAt = time.Now()
//	user.UpdateAt = time.Now()
//	if err := user.Save(c); err != nil {
//		middleware.ResponseError(c, 2007, err)
//		return
//	}
//	middleware.ResponseSuccess(c, "")
//	return
//}
//
//func (demo *Api) EditUser(c *gin.Context) {
//	editInput := &dto.EditUserInput{}
//	if err := editInput.BindingValidParams(c); err != nil {
//		middleware.ResponseError(c, 2006, err)
//		return
//	}
//
//	user := &dao.User{}
//	if userDb, err := user.Find(c, int64(editInput.Id)); err != nil {
//		middleware.ResponseError(c, 2006, err)
//		return
//	} else {
//		user = userDb
//	}
//	user.Name = editInput.Name
//	user.Sex = editInput.Sex
//	user.Age = editInput.Age
//	user.Birth = editInput.Birth
//	user.Addr = editInput.Addr
//	user.UpdateAt = time.Now()
//	if err := user.Save(c); err != nil {
//		middleware.ResponseError(c, 2007, err)
//		return
//	}
//	middleware.ResponseSuccess(c, "")
//	return
//}
//
//func (demo *Api) RemoveUser(c *gin.Context) {
//	removeInput := &dto.RemoveUserInput{}
//	if err := removeInput.BindingValidParams(c); err != nil {
//		middleware.ResponseError(c, 2006, err)
//		return
//	}
//
//	user := &dao.User{}
//	if err := user.Del(c, strings.Split(removeInput.IDS, ",")); err != nil {
//		middleware.ResponseError(c, 2007, err)
//		return
//	}
//	middleware.ResponseSuccess(c, "")
//	return
//}
