package src

import (
	"app/src/metrics"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
)

func (repository *EntitiesRepo) UploadEntitiesToDb(c *gin.Context) {
	client := NewClient(os.Getenv("URL"))
	entities := client.FetchEntities()
	
	result := repository.Db.Create(entities)
	recordsCount := result.RowsAffected
	err := result.Error

	if err != nil {
		msg := "Failed to upload Entities from " + os.Getenv("URL") + ": " + err.Error()
		sendResponseStatusInternalServerError(c, msg)
		return
	}

	msg := gin.H{"message": "Uploaded " + strconv.FormatInt(recordsCount, 10) + " entities."}
	sendResponseStatusOk(c, msg)
}

func (repository *EntitiesRepo) GetEntities(c *gin.Context) {
	var entities []Entity

	err := repository.Db.Find(&entities).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "Did not find any Entity"
			sendResponseStatusNotFound(c, msg)
			return
		}
		msg := "Failed to get Entities: " + err.Error()
		sendResponseStatusInternalServerError(c, msg)
		return
	}
	sendResponseStatusOk(c, entities)
}

func (repository *EntitiesRepo) GetEntity(c *gin.Context) {
	var entity Entity

	id := c.Param("id")
	err := repository.Db.Where("id = ?", id).First(&entity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "Did not find an Entity with id=" + id + ": " + err.Error()
			sendResponseStatusNotFound(c, msg)
			return
		}

		msg := "Failed to find Entity with id " + id + ": " + err.Error()
		sendResponseStatusInternalServerError(c, msg)
		return
	}
	sendResponseStatusOk(c, entity)
}

func (repository *EntitiesRepo) DeleteEntity(c *gin.Context) {
	var entity Entity

	id := c.Param("id")
	err := repository.Db.Where("id = ?", id).Delete(&entity).Error

	if err != nil {
		msg := "Failed to delete Entity with id " + id + ": " + err.Error()
		sendResponseStatusInternalServerError(c, msg)
		return
	}
	msg := gin.H{"message": "Entity deleted successfully."}
	sendResponseStatusOk(c, msg)
}

func (repository *EntitiesRepo) DeleteEntities(c *gin.Context) {
	err := repository.Db.Exec("DELETE FROM entities").Error

	if err != nil {
		msg := "Failed to delete all Entities: " + err.Error()
		sendResponseStatusInternalServerError(c, msg)
		return
	}
	msg := gin.H{"message": "Entities deleted successfully"}
	sendResponseStatusOk(c, msg)
}

func sendResponseStatusOk(c *gin.Context, msg any) {
	metrics.HttpRequestCountTotal.Inc()
	metrics.HttpRequestCountSuccessful.Inc()
	c.JSON(http.StatusOK, msg)
}

func sendResponseStatusInternalServerError(c *gin.Context, msg string) {
	metrics.HttpRequestCountTotal.Inc()
	metrics.HttpRequestCountError.Inc()
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": msg})
}

func sendResponseStatusNotFound(c *gin.Context, msg string) {
	metrics.HttpRequestCountTotal.Inc()
	metrics.HttpRequestCountNotFound.Inc()
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": msg})
}
