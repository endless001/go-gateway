package dao

import (
	"github.com/gin-gonic/gin"
	"go-gateway/model"
)

func (d *Dao) FindApp(c *gin.Context)(*model.App,error)  {
	model :=&model.App{}
	return  model,nil
}

func (d *Dao) SaveApp (model *model.App) error  {
	return nil
}



