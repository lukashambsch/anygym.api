package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store/datastore"
)

const StatusId = "status_id"
const InvalidStatusId = "Invalid " + StatusId

// if field maps to true then partial matches will be returned
var statusFields map[string]bool = map[string]bool{
	"status_id":   false,
	"status_name": true,
}

func BuildWhere(fields map[string]bool, params url.Values) string {
	var (
		where string = "WHERE"
		count int    = len(params)
		i     int    = 1
	)
	if count == 0 {
		return ""
	}

	for k, v := range params {
		if _, ok := fields[k]; ok {
			if fields[k] == true {
				where = fmt.Sprintf("%s %s LIKE '%%%s%%'", where, k, v[0])
			} else {
				where = fmt.Sprintf("%s %s = '%s'", where, k, v[0])
			}

			if i < count {
				where += " AND"
			}
		}
		i += 1
	}

	if where == "WHERE" {
		return "LIMIT 0"
	}

	return where
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
	where := BuildWhere(statusFields, c.Request.URL.Query())
	statuses, err := datastore.GetStatusList(where)
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
