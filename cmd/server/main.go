package main

import (
	"context"
	"fmt"
	"log"
	"main/cmd/app"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Boilerplate API docs
// @version 1.0
// @description This is a boilerplate API server.

// @contact.name API Support
// @contact.email aji.zapar00@gmail.com

// @host localhost:8080
// @BasePath /api/v1
func main() {
	serveCtx, cancelServeCtx := context.WithCancel(context.Background())

	// initialize the app
	app, err := app.NewApp(serveCtx)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sigCh

		cleanCtx, cancelCleanCtx := context.WithTimeout(serveCtx, 30*time.Second)

		go func() {
			<-cleanCtx.Done()

			if cleanCtx.Err() == context.DeadlineExceeded {
				log.Fatal("Graceful shutdown timed out... Forcing exit now...")
			}
		}()

		log.Println("Gracefully shutdown server...")

		err := app.Clean(cleanCtx)
		if err != nil {
			log.Fatal(err)
		}

		cancelCleanCtx()
		cancelServeCtx()
	}()

	// Start the server
	log.Println("Starting server...")

	err = app.Serve(serveCtx)
	fmt.Println("error ga ", err)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	<-serveCtx.Done()
}
