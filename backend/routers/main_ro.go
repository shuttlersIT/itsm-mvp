package routers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

func SetupMainRouter2(router *gin.Engine, db *sql.DB) {
	router.Use(middleware.ApiMiddleware(db))

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
	router.Use(sessions.Sessions("itsmsession", store))

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

	// Add other main routes as needed

	// Index Route Router Group
	authorized := router.Group("/")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/itsm", handlers.ItsmHandler)
		authorized.GET("/assets", handlers.ItsmHandler)
		authorized.GET("/procurement", handlers.ItsmHandler)
		authorized.GET("/admin", handlers.ItsmHandler)
		//authorized.GET("/testing", handlers.HomeTest)
	}

	// Add other main groups and routes as needed

	// Example routes for creating entities
	//router.POST("/createStaff", handlers.CreateStaffHandler)
	//router.POST("/createAgent", handlers.CreateAgentHandler)
	//router.POST("/createStatus", handlers.CreateStatusHandler)
	//router.POST("/createAsset", handlers.CreateAssetHandler)
}
