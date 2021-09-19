package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	eg, egCtx := errgroup.WithContext(ctx)

	//start server
	srv := &http.Server{Addr: ":8088"}
	eg.Go(func() error {
		return startServer(srv)
	})

	//stop server
	eg.Go(func() error {
		<-egCtx.Done() //wait for stop signal
		return srv.Shutdown(egCtx)
	})

	//detect signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	eg.Go(func() error {
		for {
			select {
			case <-egCtx.Done():
				return egCtx.Err()
			case <-ch:
				cancel()
			}
		}
		return nil
	})

	//wait
	if err := eg.Wait(); err != nil {
		fmt.Println("http server error ocur: ", err)
	}
}

func startServer(srv *http.Server) error {
	http.HandleFunc("/test1", test1Handler)
	http.HandleFunc("/test2", test2Handler)
	return srv.ListenAndServe()
}

func test1Handler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "test1")
}

func test2Handler(w http.ResponseWriter, req *http.Request) {
	panic("error 1")
	io.WriteString(w, "test2")
}
