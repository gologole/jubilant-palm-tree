package transport

import (
	"cmd/main.go/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// MyHandler handles HTTP requests
type MyHandler struct {
	service *service.Service
}

// NewMyHandler creates a new instance of MyHandler
func NewMyHandler(service *service.Service) *MyHandler {
	return &MyHandler{
		service,
	}
}

// InitRoutes initializes the routes
func (h *MyHandler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "Hello, World!")
	})

	return router
}
