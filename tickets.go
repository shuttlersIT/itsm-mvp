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

func tickets(c *gin.Context) {

	database.TableExists(db, "tickets")

	//API ROUTER
	api := gin.Default()
	api.Use(middleware.ApiMiddleware(db))

	//RedisHost := "127.0.0.1"
	//RedisPort := "6379"

	//Dashboard APP
	//router := gin.Default()
	//token, err := handlers.RandToken(64)
	//if err != nil {
	//	log.Fatal("unable to generate random token: ", err)
	//}

	token, err := handlers.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}

	store, storeError := sessions.NewRedisStore(10, "tcp", "redisDB:6379", "", []byte(token))
	if storeError != nil {
		fmt.Println(storeError)
		log.Fatal("Unable to create save session with redis session ", storeError)
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

	//

	//Index Route Router Group
	authorized := api.Group("/")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/itsm", handlers.ItsmHandler)
		authorized.GET("/assets", handlers.ItsmHandler)
		authorized.GET("/procurement", handlers.ItsmHandler)
		authorized.GET("/admin", handlers.ItsmHandler)
		authorized.GET("/testing", homeTest)
		authorized.POST("/secure")
	}

	//router.Use(static.Serve("/", static.LocalFile("./templates", true)))
	var id2 int
	ticketRouter := fmt.Sprintf("/admin/ticketing/work/%d", id2)

}
