package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

// getAlbums responds with the list of all albums as JSON.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	// fmt.Println(gin.Context)
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	albums = append(albums, newAlbum)
	c.JSON(201, gin.H{"message": newAlbum})
}
func main() {
	router := gin.Default()
	router.GET("/", getAlbums)
	router.POST("/create", postAlbums)

	router.Run(":2000")
}
