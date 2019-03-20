package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/jhontea/go-scrap/db"
)

func main() {
	// load .env
	fmt.Println("loading .env")
	godotenv.Load()
	fmt.Println("success loading .env")

	db.DBInit()

	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	serverPort := os.Getenv("PORT")
	serverString := fmt.Sprintf("run on port %s", serverPort)
	fmt.Println(serverString)

	if err := http.ListenAndServe(":"+serverPort, route); err != nil {
		log.Fatal(err)
	}
}
