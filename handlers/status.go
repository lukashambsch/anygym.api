package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
)

const StatusId = "status_id"

func GetStatus(c *gin.Context) {
	status, err := datastore.GetStatus(c.Param(StatusId))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not Found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, status)
}

func GetStatuses(c *gin.Context) {
	statuses, err := datastore.GetStatusList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting status list.",
		})
		return
	}

	c.JSON(http.StatusOK, statuses)
}

func PostStatus(c *gin.Context) {
	in := &models.Status{}
	err := c.BindJSON(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	status := models.Status{
		StatusName: in.StatusName,
	}

	created, err := datastore.CreateStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func PutStatus(c *gin.Context) {
	in := &models.Status{}
	err := c.BindJSON(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	status := models.Status{
		StatusName: in.StatusName,
	}

	updated, err := datastore.UpdateStatus(c.Param(StatusId), status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteStatus(c *gin.Context) {
	err := datastore.DeleteStatus(c.Param(StatusId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
