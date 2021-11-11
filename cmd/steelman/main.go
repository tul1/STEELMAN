package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SportJSON struct {
	ID        string `json:"id"`
	SportName string `json:"sport_name"`
	CreatedAt string `json:"created_at"`
}

var sports = []SportJSON{
	{ID: "1", SportName: "Football", CreatedAt: "Saturday, October 30, 2021"},
	{ID: "2", SportName: "Baseball", CreatedAt: "Saturday, October 31, 2021"},
}

func main() {
	dsn := "host=172.27.0.2 user=steelman password=steelman dbname=steelman_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(&Sport{SportName: "bowling"})

	router := gin.Default()
	router.GET("/sport", getSports)
	router.GET("/sport/:id", getSportsByID)
	router.POST("/sport", postSports)

	router.Run("localhost:8080")
}

func getSports(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sports)
}

type Sport struct {
	ID        string `gorm:"default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	SportName string
}

func getSportsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range sports {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sport not found"})
}

func postSports(c *gin.Context) {
	var newSport SportJSON

	if err := c.BindJSON(&newSport); err != nil {
		return
	}

	sports = append(sports, newSport)
	c.IndentedJSON(http.StatusCreated, newSport)
}
