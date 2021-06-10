package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/burhon94/clean/cmd/clean/router"
)

func main() {
	var (
		prefix         = "/api"
		addr           = ":9999"
		ctx, cancelFun = context.WithCancel(context.Background())
	)

	newMux := router.InitRouter(prefix)

	newServer := http.Server{
		Handler: newMux,
		Addr:    addr,
	}

	go func() {
		quitCh := make(chan os.Signal, 1)
		signal.Notify(quitCh, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
		sig := <-quitCh
		log.Printf("signal....%v  ✗ ", sig)
		cancelFun()

		ctx, cancelFun = context.WithTimeout(ctx, time.Second*10)
		defer cancelFun()
		err := newServer.Shutdown(ctx)
		if err != nil {
			log.Fatalf("✗ srv.Shutdown, err: %v", err)
		}
		log.Printf("server finished goroutines ✓")
	}()

	log.Printf("server started at %s ✓ ", newServer.Addr+prefix)
	err := newServer.ListenAndServe()
	if err != nil {
		log.Fatalf("server serve ✗ , err: %v", err)
	}

	log.Printf("finished main goroutine ✓")
}
