package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TheMarvelFan/GoPractice/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // this syntax is used to import a package solely for its side-effects. It will not be dirctly used in the code
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello world")

	godotenv.Load(".env") // load .env file

	port := os.Getenv("PORT") // to get values from current OS session

	if port == "" {
		log.Fatal("PORT is not set")
	}

	dbUrl := os.Getenv("DB_URL") // to get values from current OS session

	if dbUrl == "" {
		log.Fatal("Db Url is not set")
	}

	conn, errConn := sql.Open("postgres", dbUrl)

	if errConn != nil {
		log.Fatal("Cannot connect to db:", errConn)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	config := cors.Config{
		AllowOrigins:           []string{"https://*", "http://*"},
		AllowWildcard:          true,
		AllowMethods:           []string{"GET", "POST"},
		AllowHeaders:           []string{"Authorization", "Content-Type"},
		AllowCredentials:       true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             false,
		CustomSchemas:          []string{"tauri://"},
		MaxAge:                 24 * time.Hour,
		ExposeHeaders:          []string{"X-Custom-Header"},
		AllowPrivateNetwork:    true,
	}

	router := gin.Default()

	v1Router := router.Group("/v1")
	{
		v1Router.GET("/healthCheck", healthCheckRouteHandler)
		v1Router.GET("/error", errorRouteHandler)
		v1Router.POST("/users", apiCfg.createUserHandler)
	}

	router.Use(cors.New(config))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Port:", port)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// go mod vendor -> locally store dependencies in vendor folder
// similar to npm install --save in Node.js
