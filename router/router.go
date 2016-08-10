package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/lukashambsch/gym-all-over/store/datastore"
)

func Load() *gin.Engine {

	r := gin.Default()

	r.GET("/statuses", func(c *gin.Context) {

		statuses, err := datastore.GetStatusList()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error getting status list. %s", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": statuses,
		})
	})

	return r
}
