package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: 1, Title: "Lost Village", Price: 149, Quantity: 3},
	{ID: 2, Title: "Mahou Gakuen", Price: 179, Quantity: 6},
	{ID: 3, Title: "Mahou Sensou", Price: 199, Quantity: 2},
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func PostBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusOK, newBook)
}

func main() {
	r := gin.Default()

	r.GET("/", GetBooks)
	r.POST("/books", PostBook)

	r.Run(":9090")
}
