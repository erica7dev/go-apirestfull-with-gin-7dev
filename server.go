package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string     `json:"id"`
	Title  string  `json:"title"`
	Lyrics string  `json:"lyrics"`
	Year   int     `json:"year"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albuns = []album{
	{ID: "1", Title: "The Dark Side of the Moon", Lyrics: "The Dark Side of the Moon", Year: 1973, Artist: "Pink Floyd", Price: 10.99},
	{ID: "2", Title: "Thriller", Lyrics: "Billie Jean", Year: 1980, Artist: "Michael Jackson", Price: 9.99},
	{ID: "3", Title: "Celebration", Lyrics: "La Isla Bonita", Year: 1985, Artist: "Madonna", Price: 18.34},
}

func getAlbuns(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albuns)
}

func createAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add new album to the slice
	albuns = append(albuns, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albuns {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albuns {
		if a.ID == id {
			var newAlbum album
			if err := c.BindJSON(&newAlbum); err != nil {
				return
			}
			albuns[i] = newAlbum
			c.IndentedJSON(http.StatusCreated, newAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albuns {
		if a.ID == id {
			albuns = append(albuns[:i], albuns[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"Message": "Album deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found"})
}

func main (){
	router := gin.Default()
	router.GET("/albuns", getAlbuns)
	router.GET("/albuns/:id", getAlbum)
	router.POST("/albuns", createAlbum)
	router.PUT("/albuns/:id", updateAlbum)
	router.DELETE("/albuns/:id", deleteAlbum)
	router.Run(":8080")
}
