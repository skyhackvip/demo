package service

import (
	"github.com/skyhackvip/geek/sproject/configs"
	"github.com/skyhackvip/geek/sproject/internal/dao"
)

type Service struct {
	c   *configs.Config
	dao *dao.Dao
}

func New(c *configs.Config) *Service {
	return &Service{
		c:   c,
		dao: dao.New(c),
	}
}
