package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-service/handler"
	"time"
)

const PORT = 5000

func main() {

	logger := log.New(os.Stdout, "product-service", log.LstdFlags)
	mux := http.NewServeMux()

	mux.Handle("/", handler.NewProductHandler(logger))

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", PORT),
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Server started on port:", PORT)
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	logger.Println("Gracefully shutted down", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)

}
