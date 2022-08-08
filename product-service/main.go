package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-service/api/v1/handler"
	"product-service/api/v1/product"
	"product-service/db"
	"time"

	"github.com/gin-gonic/gin"
)

const PORT = "5000"

func CustomServiceLog(l *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		l.Println(c.Request.Method, c.Request.URL)
		c.Next()
	}
}

func main() {

	logger := log.New(os.Stdout, "product-service", log.LstdFlags)

	router := gin.Default()

	router.Use(CustomServiceLog(logger))

	db, err := db.Init()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(logger, productService)

	// routing
	apiV1 := router.Group("/api/v1/")
	apiV1.GET("/", productHandler.Greetings)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", PORT),
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start server
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
