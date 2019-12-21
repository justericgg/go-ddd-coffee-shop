package main

import (
	"fmt"
	"github.com/justericgg/go-ddd-coffee-shop/order/infra/api"
	"log"
	"net/http"
	"time"
)

const Port = 8888

func main() {
	router := api.Route()
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("Server listen error: %v", err)
	}
}
