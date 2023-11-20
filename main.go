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
	//initiate mysql database
	status, db := database.ConnectMysql()
	fmt.Println(status)
	database.TableExists(db, "tickets")

	//API ROUTER
	api := gin.Default()

	api.Use(middleware.ApiMiddleware(db))

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
	api.Use(sessions.Sessions("itsmsession", store))

	// We check for errors.

	//sessionStore, _ := sessions.NewRedisStore(10, "tcp", RedisHost+":"+RedisPort, "", []byte(token))
	//router.Use(sessions.Sessions("portalsession", sessionStore))

	api.Use(gin.Logger())
	api.Use(gin.Recovery())
	api.Static("/css", "templates/css")
	api.Static("/img", "templates/img")
	api.Static("/js", "templates/js")
	api.LoadHTMLGlob("templates/*.html")

	api.GET("/", handlers.IndexHandler)
	api.GET("/index", handlers.IndexHandler)
	api.GET("/login", handlers.LoginHandler)
	api.GET("/auth", handlers.AuthHandler)
	api.GET("/logout", handlers.LogoutHandler)
	api.GET("/login/admin", handlers.GetAgentHandler)

	//Login Form Handlers
	api.POST("/register", handlers.Register)
	api.POST("/formlogin", handlers.PasswordLogin)

	//Index Route Router Group
	authorized := api.Group("/")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/itsm", handlers.ItsmHandler)
		authorized.GET("/assets", handlers.ItsmHandler)
		authorized.GET("/procurement", handlers.ItsmHandler)
		authorized.GET("/admin", handlers.ItsmHandler)
		authorized.GET("/testing", homeTest)
	}

	//router.Use(static.Serve("/", static.LocalFile("./templates", true)))
	var id2 int
	ticketRouter := fmt.Sprintf("/admin/ticketing/work/%d", id2)

	//ADMIN Router Group
	admin2 := api.Group("/admin")
	admin2.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/admin/testing", homeTest)
		authorized.GET("/admin/itsm/ticketing", handlers.ItDeskAdminHandler)
		authorized.GET("/admin/assets", handlers.AssetsAdminHandler)
		authorized.GET("/admin/procurement", handlers.ProcurementAdminHandler)
		authorized.GET(ticketRouter, handlers.GetTicket2)
		//authorized.GET("/ticketing/0/admin/work/:id", handlers.GetTicket)
		//authorized.GET("/admin/ticketing/work", handlers.ListTickets)
		//authorized.POST("/admin/ticketing/work", handlers.CreateTicket)
		//authorized.PUT("/admin/ticketing/work/:id", handlers.UpdateTicket)
		//authorized.DELETE("/admin/ticketing/work/:id", handlers.DeleteTicket)
	}

	//ITSM Portal Router Group
	itsm := api.Group("/itsm")
	itsm.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/itsm/ticketing/itportal", handlers.ItDeskPortalHandler)
		authorized.GET("/itsm/ticketing", handlers.ItDeskHandler)
		authorized.GET("/itsm//testing", homeTest)
		authorized.GET(ticketRouter, handlers.GetTicket2)
		//authorized.GET("itsm/ticketing/0/admin/work/:id", handlers.GetTicket)
	}

	//ASSETS Portal Router Group
	assets := api.Group("/assets")
	assets.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/assets/portal", handlers.AssetsPortalHandler)
		authorized.GET("/assets", handlers.AssetsHandler)
		authorized.GET("/assets/testing", homeTest)
	}

	//PROCUREMENT Dashboard Portal Router Group
	procurement := api.Group("/procurement")
	procurement.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/procurement/portal", handlers.ProcurementPortalHandler)
		authorized.GET("/procurement", handlers.ProcurementHandler)
		authorized.GET("/procurement/testing", homeTest)
	}

	if err := api.Run(":5152"); err != nil {
		log.Fatal(err)
	}

	/////////////////////////////////////////////////////////////////////////////////
	//MAIN ROUTER
	router := gin.Default()

	router.Use(middleware.ApiMiddleware(db))

	//RedisHost := "127.0.0.1"
	//RedisPort := "6379"

	//Dashboard APP
	//router := gin.Default()
	token1, err := handlers.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}

	store1, storeError1 := sessions.NewRedisStore(10, "tcp", "redisDB:6379", "", []byte(token1))
	if storeError1 != nil {
		fmt.Println(storeError1)
		log.Fatal("Unable to create work with redis session ", storeError1)
	}
	//store := sessions.NewCookieStore([]byte(token))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		Secure:   true,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("itsmsession", store1))

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

	//Login Form Handlers
	router.POST("/register", handlers.Register)
	router.POST("/formlogin", handlers.PasswordLogin)

	//Index Route Router Group
	authorized1 := router.Group("/")
	authorized1.Use(middleware.AuthorizeRequest())
	{
		authorized1.GET("/itsm", handlers.ItsmHandler)
		authorized1.GET("/assets", handlers.ItsmHandler)
		authorized1.GET("/procurement", handlers.ItsmHandler)
		authorized1.GET("/admin", handlers.ItsmHandler)
		authorized1.GET("/testing", homeTest)
	}

	//router.Use(static.Serve("/", static.LocalFile("./templates", true)))
	var id int
	ticketRouter1 := fmt.Sprintf("/admin/ticketing/work/%d", id)

	//ADMIN Router Group
	admin1 := router.Group("/admin")
	admin1.Use(middleware.AuthorizeRequest())
	{
		authorized1.GET("/admin/testing", homeTest)
		authorized1.GET("/admin/itsm/ticketing", handlers.ItDeskAdminHandler)
		authorized1.GET("/admin/assets", handlers.AssetsAdminHandler)
		authorized1.GET("/admin/procurement", handlers.ProcurementAdminHandler)
		authorized1.GET(ticketRouter1, handlers.GetTicket2)
		//authorized.GET("/ticketing/0/admin/work/:id", handlers.GetTicket)
		//authorized.GET("/admin/ticketing/work", handlers.ListTickets)
		//authorized.POST("/admin/ticketing/work", handlers.CreateTicket)
		//authorized.PUT("/admin/ticketing/work/:id", handlers.UpdateTicket)
		//authorized.DELETE("/admin/ticketing/work/:id", handlers.DeleteTicket)
	}

	//ITSM Portal Router Group
	itsm1 := router.Group("/itsm")
	itsm1.Use(middleware.AuthorizeRequest())
	{
		authorized1.GET("/itsm/ticketing/itportal", handlers.ItDeskPortalHandler)
		authorized1.GET("/itsm/ticketing", handlers.ItDeskHandler)
		authorized1.GET("/itsm//testing", homeTest)
		authorized1.GET(ticketRouter1, handlers.GetTicket2)
		//authorized.GET("itsm/ticketing/0/admin/work/:id", handlers.GetTicket)
	}

	//ASSETS Portal Router Group
	assets1 := router.Group("/assets")
	assets1.Use(middleware.AuthorizeRequest())
	{
		authorized1.GET("/assets/portal", handlers.AssetsPortalHandler)
		authorized1.GET("/assets", handlers.AssetsHandler)
		authorized1.GET("/assets/testing", homeTest)
	}

	//PROCUREMENT Dashboard Portal Router Group
	procurement1 := router.Group("/procurement")
	procurement1.Use(middleware.AuthorizeRequest())
	{
		authorized1.GET("/procurement/portal", handlers.ProcurementPortalHandler)
		authorized1.GET("/procurement", handlers.ProcurementHandler)
		authorized1.GET("/procurement/testing", homeTest)
	}

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
