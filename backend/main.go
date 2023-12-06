package main

import (
	"fmt"
	"log"

	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/database"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/routers"
)

func main() {

	//initiate mysql database
	status, db := database.ConnectMysql()
	fmt.Println(status)
	database.TableExists(db, "tickets")

	// API Router
	api := gin.Default()
	routers.SetupAPIRouter(api, db)

	// Main Router
	mainRouter := gin.Default()
	routers.SetupMainRouter(mainRouter, db)

	// Example routes for creating entities
	mainRouter.POST("/createStaff", handlers.CreateStaffHandler)
	mainRouter.POST("/createAgent", handlers.CreateAgentHandler)
	mainRouter.POST("/createStatus", handlers.CreateStatusHandler)
	mainRouter.POST("/createAsset", handlers.CreateAssetHandler)

	// Run API and Main Routers
	go func() {
		if err := api.Run(":5152"); err != nil {
			log.Fatal(err)
		}
	}()

	if err := mainRouter.Run(":5151"); err != nil {
		log.Fatal(err)
	}
}

func HomeTest1(c *gin.Context) {
	s := sessions.Default(c)
	var count int
	v := s.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count += 1
	}
	s.Set("count", count)
	s.Save()

	c.JSON(200, gin.H{"count": count})
}

func HomeTest2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from homeTest!",
	})
}
