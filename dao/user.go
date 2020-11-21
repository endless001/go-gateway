package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-gateway/ent"
	"go-gateway/ent/user"
)

func (d *Dao) CheckUser(c *gin.Context, userName string, password string) (*ent.User, error) {
	user, err := d.client.User.Query().
		Where(user.UsernameEQ(userName), user.IsDeleteEQ(0)).
		Only(c)
	if err != nil {

	}
}
func (d *Dao) FindUser(c context.Context) (*ent.User, error) {
	return nil, nil
}
