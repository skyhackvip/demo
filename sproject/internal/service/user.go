package service

import (
	"context"
	"errors"
	"github.com/skyhackvip/geek/sproject/api"
	"github.com/skyhackvip/geek/sproject/configs"
	"log"
)

func (s *Service) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	isLogined, err := s.dao.CheckLogin(req.Name, req.Pass)
	var msg string
	if err != nil {
		log.Println("user service:", err)
		err = errors.New(configs.SqlError)
	}
	return &api.LoginResponse{IsLogined: isLogined, Msg: msg}, err
}
