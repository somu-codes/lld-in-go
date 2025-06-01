package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"url-shortener/internal/api"
	"url-shortener/internal/service"
	"url-shortener/internal/store"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	urlStore := store.NewInMemoryStore(make(map[string]string), make(map[string]string))
	urlService := service.NewURLService(urlStore)
	handler := api.NewHandler(urlService)

	router := gin.Default()

	router.POST("/shorten", handler.HandleShorten)
	router.GET("/:code", handler.HandleResolve)

	log.Println("Starting server at http://localhost:8080")
	router.Run(":8080")
}
