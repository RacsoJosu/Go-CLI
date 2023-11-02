package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RacsoJosu/Api-Rest/server"
)

func main() {
	ctx := context.Background()
	// crear un channel
	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":3000")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	log.Println("Server started")
	<- serverDoneChan
	srv.Shutdown(ctx)

	
	log.Println("Server stopped")
	

}
