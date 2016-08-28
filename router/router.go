package router

import (
    "fmt"

	"github.com/gin-gonic/gin"

	"github.com/lukashambsch/gym-all-over/handlers"
)

var V1URLBase string = "/api/v1"

func Load() *gin.Engine {

	r := gin.Default()

    status := r.Group(fmt.Sprintf("%s%s", V1URLBase, "/statuses"))
    {
        status.GET("", handlers.GetStatuses)
        status.GET("/:status_id", handlers.GetStatus)
        status.POST("", handlers.PostStatus)
        status.PUT("/:status_id", handlers.PutStatus)
        status.DELETE("/:status_id", handlers.DeleteStatus)
    }

	return r
}
