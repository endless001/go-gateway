package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/internal/database"
	"go-gateway/internal/ent"
)

type TcpRuleDao struct {
}

func (d *TcpRuleDao) GetTcpRule(c *gin.Context, id int64) (*ent.TcpRule, error) {
	model, err := database.Client.TcpRule.
		Get(c, id)
	return model, err
}

func (d *TcpRuleDao) CreateTcpRule(c *gin.Context, rule *ent.TcpRule) error {
	_, err := database.Client.TcpRule.Create().
		SetID(rule.ID).
		Save(c)
	return err
}
