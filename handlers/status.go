package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
)

const StatusId = "status_id"
const InvalidStatusId = "Invalid " + StatusId

var statusFields map[string]string = map[string]string{
	"status_id":   "int",
	"status_name": "string",
}

func GetStatus(c *gin.Context) {
	statusId, err := strconv.ParseInt(c.Param(StatusId), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": InvalidStatusId,
		})
		return
	}

	status, err := datastore.GetStatus(statusId)
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
	var statement string
	query := c.Request.URL.Query()
	where, err := BuildWhere(statusFields, query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	sort, err := BuildSort(statusFields, query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	statement = fmt.Sprintf("%s %s", where, sort)

	fmt.Println(statement)
	statuses, err := datastore.GetStatusList(statement)
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
	statusId, err := strconv.ParseInt(c.Param(StatusId), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": InvalidStatusId,
		})
		return
	}

	in := &models.Status{}
	err = c.BindJSON(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	status := models.Status{
		StatusName: in.StatusName,
	}

	updated, err := datastore.UpdateStatus(statusId, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteStatus(c *gin.Context) {
	statusId, err := strconv.ParseInt(c.Param(StatusId), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": InvalidStatusId,
		})
		return
	}

	err = datastore.DeleteStatus(statusId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
