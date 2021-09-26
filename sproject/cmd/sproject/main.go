package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/skyhackvip/geek/sproject/api"
	"github.com/skyhackvip/geek/sproject/configs"
	"github.com/skyhackvip/geek/sproject/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	_confName string
)

func init() {
	flag.StringVar(&_confName, "config_name", "config.yaml", "default config filename")
}

func main() {
	//config
	flag.Parse()
	c := configs.LoadConfig(_confName)

	//grpc
	port := 7777
	listener, err := net.Listen("tcp", ftm.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}
	svc := service.New(c)
	s := grpc.NewServer()
	api.RegisterUserServer(s, svc)
	s.Serve(listener)

	//graceful restart
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//close

	select {
	case <-ctx.Done():
		log.Println("server stop timeout 5 seconds")
	}
	log.Println("server stop")
}