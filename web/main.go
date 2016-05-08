package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {

	initEnv()
	defer initLogger().Close()

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.LoadHTMLGlob("templates/*")
	router.Static("img", "./public/img")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Docker Meetup",
		})
	})

	v1 := router.Group("v1")
	{
		v1.GET("/health_check", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "healthy",
			})
		})
	}

	port := os.Getenv("HTTP_PORT")

	fmt.Println("Port: ", port)
	router.Run(fmt.Sprintf(":%s", port))
}

func initEnv() {
	env := godotenv.Load()
	if env != nil {
		log.Println("Error loading .env file")
	}
}

func initDB() (*gorm.DB, error) {
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True", dbUser, dbPass, dbHost, dbName)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	db.DB().Ping()
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(50)

	logMode := true
	if os.Getenv("GIN_MODE") == "release" {
		logMode = false
	}
	db.LogMode(logMode)

	return db, nil
}

func initLogger() *os.File {

	var logFile *os.File
	var err error
	if _, err = os.Stat("log"); os.IsNotExist(err) {
		logFile, err = os.Create("log")
		log.Println("Logfile created")
	} else {
		logFile, err = os.OpenFile("log", os.O_WRONLY, 0666)
		log.Println("Logfile opened")
	}

	if err != nil {
		log.Fatal("Logfile error: ", err)
	}

	log.SetOutput(logFile)

	return logFile
}

// CORSMiddleware test
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, If-Modified-Since, expiry, uid, token-type, access-token, client, accept")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
