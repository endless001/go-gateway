package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/dao"
	"go-gateway/internal/ent"
	"go-gateway/internal/util"
)

type UserService struct {
	ud *dao.UserDao
}

func (s *UserService) CheckUser(c *gin.Context, userName, password string) (*ent.User, error) {
	s.ud = new(dao.UserDao)

	user, err := s.ud.CheckUser(c, userName)

	if err != nil {
		return nil, errors.New("用户名不存在!")
	}
	saltPassword := util.GenSaltPassword(user.Salt, password)

	if user.Password != saltPassword {
		return nil, errors.New("密码错误请重新输入!")
	}

	return user, nil
}

func (s *UserService) FindUser(c *gin.Context, userName string) (*ent.User, error) {
	s.ud = new(dao.UserDao)

	user, err := s.ud.CheckUser(c, userName)

	if err != nil {
		return nil, errors.New("用户名不存在!")
	}

	return user, nil
}

func (s *UserService) ChangePassword(c *gin.Context, userId int, password string) error {
	s.ud = new(dao.UserDao)

	user := (&ent.User{ID: userId, Password: password})

	err := s.ud.UpdateUser(c, user)

	if err != nil {
		return errors.New("修改密码错误!")
	}

	return nil
}
