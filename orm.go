package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Title string
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Device{})

	router := gin.Default()
	router.GET("/devices", getDevices)
	router.GET("/devices/:id", deviceById)
	router.POST("/devices", createDevice)
	router.PATCH("/devices/:id", updateDevice)

	router.Run("localhost:8088")

}

func deviceById(c *gin.Context) {
	id := c.Param("id")
	book, err := getDeviceById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getDeviceById(id string) (*Device, error) {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var device Device
	result := db.First(&device, "ID = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Device not found")
		}
		return nil, result.Error
	}

	return &device, nil
}

func createDevice(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var newDevice device

	if err := c.BindJSON(&newDevice); err != nil {
		fmt.Print(err)
		return
	}
	var _device = db.Create(&Device{Title: newDevice.Title})

	c.IndentedJSON(http.StatusCreated, _device)
}

func getDevices(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	var devices []Device
	db.Find(&devices)

	c.IndentedJSON(http.StatusOK, devices)
}

func updateDevice(c *gin.Context) {

}

type device struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
