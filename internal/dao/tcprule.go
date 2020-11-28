package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/tcprule"
)

func (d *Dao) FindTcpRule(c *gin.Context, id int64) (*ent.TcpRule, error) {
	model, err := d.client.TcpRule.Query().
		Where(tcprule.ID(id)).
		Only(c)
	return model, err
}

func (d *Dao) SaveTcpRule(c *gin.Context, rule ent.TcpRule) error {
	_, err := d.client.TcpRule.Create().
		SetID(rule.ID).
		Save(c)
	return err
}

func (d *Dao) GetTcpRule(c *gin.Context, serviceId int64) ([]*ent.TcpRule, int, error) {
	list, err := d.client.TcpRule.Query().
		Where(tcprule.ServiceIDEQ(serviceId)).
		All(c)
	if err != nil {
		return nil, 0, err
	}

	count := len(list)
	return list, count, nil
}
