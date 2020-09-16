package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Neil-Uli/Restful-go/railAPI/dbutils"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID              int    `json:"id"`
	DriverName      string `json:"driverName"`
	OperatingStatus bool   `json:"operatingStatus"`
}

// StationResource holds information about locations
type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

// GetTrain returns the train detail
func GetTrain(c *gin.Context) {
	var train TrainResource
	id := c.Param("train_id")
	err := DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS from train where id=?", id).Scan(&train.ID, &train.DriverName, &train.OperatingStatus)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": train,
		})
	}
}

// GetStation returns the station detail
func GetStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station_id")
	err := DB.QueryRow("select ID, Name, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station where id=?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}
}

// CreateStation handles the POST
func CreateStation(c *gin.Context) {
	var station StationResource
	// Parse the body into our resrource
	if err := c.BindJSON(&station); err == nil {
		// Format Time to Go time format
		statement, _ := DB.Prepare("insert into station (NAME, OPENING_TIME, CLOSING_TIME) values (?, ?, ?)")
		result, _ := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err == nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

// RemoveStation handles the removing of resource
func RemoveStation(c *gin.Context) {
	id := c.Param("station-id")
	statement, _ := DB.Prepare("delete from station where id=?")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}
func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(DB)
	r := gin.Default()
	// Add routes to REST verbs
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station_id", RemoveStation)
	r.GET("/v1/trains/:train_id", GetTrain)
	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
