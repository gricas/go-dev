// package comment
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
/*
The tutorial includes the following sections:

    Design API endpoints.
    Create a folder for your code.
    Create the data.
    Write a handler to return all items.
    Write a handler to add a new item.
    Write a handler to return a specific item.

*/

// album represents data about a record album
type album struct {
	ID 		string  `json:"id"`
	Title	string  `json:"title"`
	Artist  string  `json:"artist"`
	Price	float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gery Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}