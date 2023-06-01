package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type home struct {
	Message string `json:"message"`
}

var albums = []album{
	{ID: "1", Title: "Blue train", Artist: "John Coltrane", Price: 56.99},
}

func getAlbums(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, albums)
}
func homeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, home{Message: "Hello World!!"})
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/", homeHandler)
	router.Run()
	// fmt.Println("Hello world")
}
