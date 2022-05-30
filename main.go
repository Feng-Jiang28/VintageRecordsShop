package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Write a handler to return all items
// 1. Logic to prepare a response
// 2. Code to map the request path to your logic

func main() {
	router := gin.Default()

	//To associate the GET http method and /albums with a handler function
	router.GET("/albums", getAlbums) // passing the function

	router.Run("localhost:8080")
}

// gin.Context request details, validates and serializes JSON,
// Context.IndentedJSON serialize the struct into Json and added it to the response.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// go get .
// add dependency for your module. . means get dependencies for code in the current directory.