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
	"github.com/shuttlersIT/itsm-mvp/handlers"
	"github.com/shuttlersIT/itsm-mvp/middleware"
)

func main() {
	router := gin.Default()

	//initiate mysql database
	status, db := database.ConnectMysql()
	fmt.Println(status)
	database.TableExists(db, "tickets")

	router.Use(middleware.ApiMiddleware(db))

	//RedisHost := "127.0.0.1"
	//RedisPort := "6379"

	//Dashboard APP
	//router := gin.Default()
	token, err := handlers.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}

	store, storeError := sessions.NewRedisStore(10, "tcp", "redisDB:6379", "", []byte(token))
	if storeError != nil {
		fmt.Println(storeError)
		log.Fatal("Unable to create work with redis session ", storeError)
	}
	//store := sessions.NewCookieStore([]byte(token))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		Secure:   true,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("portalsession", store))

	// We check for errors.

	//sessionStore, _ := sessions.NewRedisStore(10, "tcp", RedisHost+":"+RedisPort, "", []byte(token))
	//router.Use(sessions.Sessions("portalsession", sessionStore))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/css", "templates/css")
	router.Static("/img", "templates/img")
	router.Static("/js", "templates/js")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", handlers.IndexHandler)
	router.GET("/index", handlers.IndexHandler)
	router.GET("/login", handlers.LoginHandler)
	router.GET("/auth", handlers.AuthHandler)
	router.GET("/logout", handlers.LogoutHandler)
	router.GET("/login/admin", handlers.GetAgentHandler)

	//Index Route Router Group
	authorized := router.Group("/")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/itsm", handlers.ItsmHandler)
		authorized.GET("/testing", homeTest)
	}
	//router.Use(static.Serve("/", static.LocalFile("./templates", true)))
	var id int
	ticketRoute := fmt.Sprintf("/ticketing/0/admin/work/%d", id)

	//ITSM Portal Router Group
	itsm := router.Group("/itsm/ticketing")
	itsm.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/itsm/ticketing/itportal", handlers.ItDeskPortalHandler)
		authorized.GET("/itsm/ticketing/0/admin", handlers.ItDeskAdminHandler)
		authorized.GET("/itsm/ticketing", handlers.ItDeskHandler)
		authorized.GET("/testing", homeTest)
		authorized.GET(ticketRoute, handlers.GetTicket)
		//authorized.GET("/ticketing/0/admin/work/:id", handlers.GetTicket)
		authorized.GET("/ticketing/0/admin/work", handlers.ListTickets)
		authorized.POST("/ticketing/0/admin/work", handlers.CreateTicket)
		authorized.PUT("/ticketing/0/admin/work/:id", handlers.UpdateTicket)
		authorized.DELETE("/ticketing/0/admin/work/:id", handlers.DeleteTicket)

	}

	//Assets Portal Router Group
	assets := router.Group("/itsm/assets")
	assets.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/itsm/assets/portal", handlers.AssetsPortalHandler)
		authorized.GET("/itsm/assets/1/admin", handlers.AssetsAdminHandler)
		authorized.GET("/itsm/assets", handlers.AssetsHandler)
		authorized.GET("/testing", homeTest)
	}

	//Procurement Dashboard Portal Router Group
	procurement := router.Group("/itsm/procurement")
	procurement.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/itsm/procurement/portal", handlers.ProcurementPortalHandler)
		authorized.GET("/itsm/procurement/2/admin", handlers.ProcurementAdminHandler)
		authorized.GET("/itsm/procurement", handlers.ProcurementHandler)
		authorized.GET("/testing", homeTest)
	}

	if err := router.Run(":5152"); err != nil {
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
