package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
	"go-gateway/internal/ent/tcprule"
)

type TcpRuleDao struct {
}

func (d *TcpRuleDao) FindTcpRule(c *gin.Context, id int64) (*ent.TcpRule, error) {
	model, err := database.Client.TcpRule.Query().
		Where(tcprule.ID(id)).
		Only(c)
	return model, err
}

func (d *TcpRuleDao) SaveTcpRule(c *gin.Context, rule ent.TcpRule) error {
	_, err := database.Client.TcpRule.Create().
		SetID(rule.ID).
		Save(c)
	return err
}

func (d *TcpRuleDao) GetTcpRule(c *gin.Context, serviceId int64) ([]*ent.TcpRule, int, error) {
	list, err := database.Client.TcpRule.Query().
		Where(tcprule.ServiceIDEQ(serviceId)).
		All(c)
	if err != nil {
		return nil, 0, err
	}

	count := len(list)
	return list, count, nil
}
