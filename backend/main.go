package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"math/rand"

	"github.com/gin-contrib/cors"
)

type ShortlyLink struct{
	gorm.Model
	OriginalURL string `gorm:"unique"`
	ShortURL string `gorm:"unique"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	

	dsn := dbUsername + ":" + dbPassword + "@tcp(" +  dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&ShortlyLink{})

	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/shorten", func(c *gin.Context) {
		var data struct {
			URL string `json:"url" binding:"required"`
		}

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		var link ShortlyLink
		result := db.Where("original_url = ?", data.URL).First(&link)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound{
				shortURL := generateShortURL()
				link = ShortlyLink{OriginalURL: data.URL, ShortURL: shortURL}
				result = db.Create(&link)
				if result.Error != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
					return
				}
			}
		}
		c.JSON(http.StatusOK, gin.H{"short_url": link.ShortURL})
	})

	r.GET("/:shortURL", func (c *gin.Context)  {
		shortURL := c.Param("shortURL")
		var link ShortlyLink

		result := db.Where("short_url = ?", shortURL).Find(&link)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error":"URL not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			}
			return
		}
		c.Redirect(http.StatusMovedPermanently, link.OriginalURL)
	})

	r.Run(":8000")
}

func generateShortURL() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var shortURL string
	for i := 0; i < 6; i++ {
		shortURL += string(chars[rand.Intn(len(chars))])
	}

	return shortURL
}