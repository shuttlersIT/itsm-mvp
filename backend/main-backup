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
	"github.com/shuttlersIT/itsm-mvp/database"
	"github.com/shuttlersIT/itsm-mvp/routers"
)

var api gin.Engine

func main() {

	//initiate mysql database
	status, db := database.ConnectMysql()
	fmt.Println(status)
	database.TableExists(db, "tickets")

	// API router
	api := gin.Default()
	routers.SetupAPIRouter(api, db)
	if err := api.Run(":5152"); err != nil {
		log.Fatal(err)
	}

	// Main router
	router := gin.Default()
	routers.SetupMainRouter(router, db)
	if err := router.Run(":5151"); err != nil {
		log.Fatal(err)
	}

}

func homeTest(c *gin.Context) {
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

func HomeTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from homeTest!",
	})
}
