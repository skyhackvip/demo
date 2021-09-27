package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/skyhackvip/geek/sproject/api"
	"github.com/skyhackvip/geek/sproject/configs"
	"github.com/skyhackvip/geek/sproject/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	flag.StringVar(&_confName, "c", "config.yaml", "default config filename")
}

func main() {
	//config
	flag.Parse()
	conf, err := configs.LoadConfig(_confName)
	if err != nil {
		log.Fatal("config error")
	}

	//rpc
	svc := service.New(conf)
	rpc := grpc.NewServer()
	go func() {

		listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.Server))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("server start")
		api.RegisterUserServer(rpc, svc)
		reflection.Register(rpc)
		rpc.Serve(listener)
	}()

	//graceful restart
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("server stop")
	rpc.GracefulStop()
	select {
	case <-ctx.Done():
		log.Println("server stop timeout 5 seconds")
	}

}
