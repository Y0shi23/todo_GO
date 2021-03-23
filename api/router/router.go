package router

import (
	"Y0shi23/todo/models"
	"Y0shi23/todo/pkg/compute"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	ctx := context.Background()
	r := gin.Default()

	// Global middleware
	r.Use(gin.Logger())

	db := compute.ConnectSQL()

	r.LoadHTMLFiles("./web/html/index.html")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/todo", func(c *gin.Context) {
		todos, _ := models.Todos().All(ctx, db)
		// if err != nil {
		// 	fmt.Errorf("Get todo error: %v", err)
		// }

		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"todo": todos,
		})
	})

	return r
}
