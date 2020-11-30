package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/user"
	"go-gateway/internal/util"
)

func (d *Dao) CheckUser(c *gin.Context, userName string, password string) (*ent.User, error) {
	user, err := d.client.User.Query().
		Where(user.UsernameEQ(userName), user.IsDeleteEQ(0)).
		Only(c)

	if err != nil {
		return nil, errors.New("用户名不存在!")
	}
	saltPassword := util.GenSaltPassword(user.Salt, password)

	if user.Password != saltPassword {
		return nil, errors.New("密码错误请重新输入!")
	}

	return user, nil
}
func (d *Dao) FindUser(c *gin.Context, id int) (*ent.User, error) {
	model, err := d.client.User.Query().
		Where(user.IDEQ(id)).
		Only(c)
	return model, err
}

func (d *Dao) SaveUser(c *gin.Context, user ent.User) error {
	_, err := d.client.User.Create().
		SetID(user.ID).
		SetCreateAt(user.CreateAt).
		Save(c)
	return err
}
