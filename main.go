package main

import (
	"chap3-challenge2/database"
	"chap3-challenge2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	database.StartDB()
	db := database.GetDB()

	routes.Routes(g, db)

	g.Run(":8080")
}
