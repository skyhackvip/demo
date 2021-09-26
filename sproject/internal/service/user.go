package service

import (
	"context"
)

func (s *Service) Login(ctx context.Context, name, pass string) bool {
	return s.dao.CheckLogin(name, pass)
}
