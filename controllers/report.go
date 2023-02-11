package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"report-generator/database"
	filegenerators "report-generator/file-generators"
	models "report-generator/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	return &UserRepo{Db: db}
}

// get visits
func (repository *UserRepo) GetVisits(c *gin.Context) {
	var visits []models.VisitVrach
	err := models.GetVisitsVrach(repository.Db, &visits)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, visits)
}

func (repository *UserRepo) GenerateExcel(c *gin.Context) {

	var visits []models.VisitVrach
	err := models.GetVisitsVrachAll(repository.Db, &visits)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	f, closeFn := filegenerators.ServeExcel(visits)
	defer closeFn()
	var b bytes.Buffer
	if err := f.Write(&b); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%d\"", &downloadName))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet; charset=utf-8")
	c.Header("Content-Description", "File Transfer")
	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}
