package service

import (
	"context"
	"errors"
	"github.com/skyhackvip/geek/sproject/api"
	"github.com/skyhackvip/geek/sproject/configs"
	"log"
	"time"
)

func (s *Service) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	log.Println("request:", req.Name, req.Pass)
	isLogined, err := s.dao.CheckLogin(req.Name, req.Pass)
	var msg string
	if err != nil {
		log.Println("user service:", err)
		err = errors.New(configs.SqlError)
	}
	if isLogined {
		msg = configs.LoginSuccess
	} else {
		msg = configs.LoginFail
	}
	time.Sleep(10 * time.Second)
	return &api.LoginResponse{IsLogined: isLogined, Msg: msg}, err
}
