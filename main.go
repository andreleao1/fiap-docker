package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func main() {
	dsn := "host=postgres user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	var db *gorm.DB
	var err error

	// Tentar conectar ao banco de dados com várias tentativas
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("Failed to connect to database. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "João"})
	db.Create(&User{Name: "Maria"})
	db.Create(&User{Name: "José"})

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		var users []User
		result := db.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})
	r.Run()
}
