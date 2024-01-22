package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PasswordRequest struct {
	InitPassword string `json:"init_password"`
}

var db *gorm.DB

type Log struct {
	gorm.Model
	Request  string
	Response string
}

func calculateNumOfSteps(password string) int {

	ans := 0
	if len(password) > 20 {
		idx := 0

		hasLower := 0
		hasUpper := 0
		hasNumber := 0

		removes := 0
		repeatedCount := 0
		targetChar := password[0]
		for idx < len(password) && idx-removes < 20 {
			if password[idx] == targetChar {
				repeatedCount++
				if repeatedCount == 3 {
					removes++
					repeatedCount = 0
				}
			} else {
				targetChar = password[idx]
				repeatedCount = 1
			}
			if repeatedCount != 0 {
				if password[idx] >= 'a' && password[idx] <= 'z' {
					hasLower = 1
				} else if password[idx] >= 'A' && password[idx] <= 'Z' {
					hasUpper = 1
				} else if password[idx] >= '0' && password[idx] <= '9' {
					hasNumber = 1
				}
			}
			idx++
		}
		removes += len(password) - idx
		adds := 3 - hasLower - hasUpper - hasNumber
		ans = removes + adds
	} else {
		idx := 0

		hasLower := 0
		hasUpper := 0
		hasNumber := 0

		replaces := 0
		repeatedCount := 0
		targetChar := password[0]
		for idx < len(password) {
			if password[idx] == targetChar {
				repeatedCount++
				if repeatedCount == 3 {
					replaces++
					repeatedCount = 0
				}
			} else {
				targetChar = password[idx]
				repeatedCount = 1
			}
			if repeatedCount != 0 {
				if password[idx] >= 'a' && password[idx] <= 'z' {
					hasLower = 1
				} else if password[idx] >= 'A' && password[idx] <= 'Z' {
					hasUpper = 1
				} else if password[idx] >= '0' && password[idx] <= '9' {
					hasNumber = 1
				}
			}
			idx++
		}
		adds := 3 - hasLower - hasUpper - hasNumber
		ans += replaces
		if len(password) < 6 {
			ans += 6 - len(password)
		}
		if adds > ans {
			ans = adds
		}
	}

	return ans
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	var err error
	db, err = gorm.Open("postgres", "host="+os.Getenv("POSTGRES_HOST")+" user="+os.Getenv("POSTGRES_USER")+" dbname="+os.Getenv("POSTGRES_DB")+" sslmode=disable password="+os.Getenv("POSTGRES_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Log{})

	r.POST("/api/strong_password_steps", func(c *gin.Context) {
		var req PasswordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		numOfSteps := calculateNumOfSteps(req.InitPassword)

		logEntry := Log{Request: req.InitPassword, Response: fmt.Sprintf(`{"num_of_steps": %d}`, numOfSteps)}
		db.Create(&logEntry)

		c.JSON(http.StatusOK, gin.H{"num_of_steps": numOfSteps})
	})

	r.Run(":8080")
}
