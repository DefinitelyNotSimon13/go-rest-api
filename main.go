package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/DefinitelyNotSimon13/go-rest-api/pb"
)

type TodoItem struct {
	ID          int32     `json:"id"`
	Uuid        uuid.UUID ``
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Completed   bool      `json:"completed"`
	Deadline    time.Time `json:"deadline,omitempty"`
}

var grpcClient pb.TodoHandlerClient

const (
	grpcServerAddr = "localhost:50051"
)

func main() {

	// REST API
	router := gin.Default()
	router.GET("/items", listItems)
	router.GET("/items/:id", getItemById)

	if err := router.Run("localhost:8090"); err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}

func init() {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(grpcServerAddr, opts...)
	if err != nil {
		log.Fatalln("fail to dial:", err)
	}

	grpcClient = pb.NewTodoHandlerClient(conn)
}

func listItems(c *gin.Context) {
	var items []*pb.TodoItem

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := grpcClient.ListTodoItems(ctx, &pb.ListTodoRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for {
		item, err := stream.Recv()
		fmt.Println("Received item...")
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		items = append(items, item.Item)
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func getItemById(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.ParseInt(idStr, 10, 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id := int32(idInt)

	item, err := grpcClient.GetTodoItem(context.Background(), &pb.ItemByIdRequest{
		Id: id,
	})

	c.IndentedJSON(http.StatusOK, item.Item)
}
