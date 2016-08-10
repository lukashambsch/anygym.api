package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/lukashambsch/gym-all-over/db"
	"github.com/lukashambsch/gym-all-over/models"
)

func Load() *gin.Engine {

	db := db.Open()
	r := gin.Default()

	var statuses models.Statuses
	status := models.Status{
		StatusName: "Pending",
	}
	statuses = append(statuses, status)

	r.GET("/statuses", func(c *gin.Context) {
		var (
			status   models.Status
			statuses models.Statuses
		)

		rows, err := db.Query("SELECT * FROM statuses;")
		if err != nil {
			log.Panic(err)
		}
		for rows.Next() {
			err = rows.Scan(&status.StatusId, &status.StatusName)
			statuses = append(statuses, status)
			if err != nil {
				log.Panic(err)
			}
		}
		defer rows.Close()

		c.JSON(http.StatusOK, gin.H{
			"result": statuses,
		})
	})

	return r
}
