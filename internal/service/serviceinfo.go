package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-gateway/internal/dao"
	"go-gateway/internal/ent"
)

type ServiceInfoService struct {
	sd *dao.ServiceInfoDao
}

func (s *ServiceInfoService) GetServiceInfo(c *gin.Context, id int64) (*ent.ServiceInfo, error) {
	model, err := s.sd.GetServiceInfo(c, id)
	if err != nil {
		return nil, errors.New("查询失败!")
	}
	return model, err
}

func (s *ServiceInfoService) DeleteTenant(c *gin.Context, id int64) error {
	err := s.sd.DeleteServiceInfo(c, id)
	if err != nil {
		return errors.New("删除失败!")
	}
	return nil
}

func (s *ServiceInfoService) CreateTenant(c *gin.Context, model *ent.ServiceInfo) error {
	err := s.sd.CreateServiceInfo(c, model)
	if err != nil {
		return errors.New("添加失败!")
	}
	return err

}

func (s *ServiceInfoService) UpdateTenant(c *gin.Context, model *ent.ServiceInfo) error {
	err := s.sd.UpdateServiceInfo(c, model)
	if err != nil {
		return errors.New("修改失败!")
	}
	return err

}
