package routers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

func SetupAPIRouter2(api *gin.Engine, db *sql.DB) {
	api.Use(middleware.ApiMiddleware(db))

	token, err := handlers.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}

	store, storeError := sessions.NewRedisStore(10, "tcp", "redisDB:6379", "", []byte(token))
	if storeError != nil {
		fmt.Println(storeError)
		log.Fatal("Unable to create save session with redis session ", storeError)
	}
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		Secure:   true,
		HttpOnly: true,
	})
	api.Use(sessions.Sessions("itsmsession", store))

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

	// Add other API routes as needed

	// Index Route Router Group
	authorized := api.Group("/")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/itsm", handlers.ItsmHandler)
		authorized.GET("/assets", handlers.ItsmHandler)
		authorized.GET("/procurement", handlers.ItsmHandler)
		authorized.GET("/admin", handlers.ItsmHandler)
		//authorized.GET("/testing", handlers.HomeTest)
	}

	// Add other API groups and routes as needed
}
