package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/user"
)

type UserDao struct {
}

func (d *UserDao) CheckUser(c *gin.Context, userName string) (*ent.User, error) {
	user, err := database.Client.User.Query().
		Where(user.UserNameEQ(userName), user.IsDeleteEQ(0)).
		Only(c)

	return user, err
}
func (d *UserDao) FindUser(c *gin.Context, id int) (*ent.User, error) {
	model, err := database.Client.User.Query().
		Where(user.IDEQ(id)).
		Only(c)
	return model, err
}

func (d *UserDao) SaveUser(c *gin.Context, user ent.User) error {
	_, err := database.Client.User.Create().
		SetID(user.ID).
		SetCreateAt(user.CreateAt).
		Save(c)
	return err
}
func (d *UserDao) UpdateUser(c *gin.Context, user *ent.User) error {
	_, err := database.Client.User.Update().
		SetPassword(user.Password).
		Save(c)
	return err
}
