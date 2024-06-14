package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server-go/api"
)

func main() {
	db := api.NewInMemoryEventStore()
	svc := api.NewEventService(db)

	r := gin.Default()

	api.RegisterHandlers(r, svc)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8084",
	}

	log.Fatal(srv.ListenAndServe())
}
