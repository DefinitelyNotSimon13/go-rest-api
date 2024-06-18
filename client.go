package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"google.golang.org/grpc"
	"pb"
)

type TodoItem struct {
	ID          int32     `json:"id"`
	Uuid        uuid.UUID ``
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Completed   bool      `json:"completed"`
	Deadline    time.Time `json:"deadline,omitempty"`
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)

	router.Run("localhost:8090")
}

// example items for testing
var items = []TodoItem{
	{
		ID:          1,
		Uuid:        uuid.New(),
		Title:       "Buy milk",
		Description: "Milk is the best",
		Completed:   false,
	},
	{
		ID:          2,
		Uuid:        uuid.New(),
		Title:       "Buy bread",
		Description: "Bread is the best",
		Completed:   false,
	},
	{
		ID:          3,
		Uuid:        uuid.New(),
		Title:       "Buy eggs",
		Description: "Eggs are the best",
		Completed:   false,
		Deadline:    time.Now().Add(time.Hour * 24 * 7),
	},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}
