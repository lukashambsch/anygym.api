package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/lukashambsch/gym-all-over/handlers"
)

func Load() *gin.Engine {

	r := gin.Default()

	r.GET("/statuses", handlers.GetStatuses)
	r.GET("/statuses/:status_id", handlers.GetStatus)
	r.POST("/statuses", handlers.PostStatus)
	r.PUT("/statuses/:status_id", handlers.PutStatus)
	r.DELETE("/statuses/:status_id", handlers.DeleteStatus)

	return r
}
