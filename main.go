package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var items = []Item{
    {ID: "1", Name: "Item One"},
    {ID: "2", Name: "Item Two"},
}

func main() {
    router := gin.Default()

	// Route ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

    // GET method
    router.GET("/items", func(c *gin.Context) {
        c.JSON(http.StatusOK, items)
    })

    // POST method
    router.POST("/items", func(c *gin.Context) {
        var newItem Item
        if err := c.ShouldBindJSON(&newItem); err == nil {
            items = append(items, newItem)
            c.JSON(http.StatusCreated, newItem)
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    // PUT method
    router.PUT("/items/:id", func(c *gin.Context) {
        id := c.Param("id")
        var updatedItem Item
        if err := c.ShouldBindJSON(&updatedItem); err == nil {
            for i, item := range items {
                if item.ID == id {
                    items[i].Name = updatedItem.Name
                    c.JSON(http.StatusOK, items[i])
                    return
                }
            }
            c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
    })

    // DELETE method
    router.DELETE("/items/:id", func(c *gin.Context) {
        id := c.Param("id")
        for i, item := range items {
            if item.ID == id {
                items = append(items[:i], items[i+1:]...)
                c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
    })

    router.Run(":5080")
}