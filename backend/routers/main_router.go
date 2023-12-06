package routers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

// SetupMainRouter sets up the main router
func SetupMainRouter(router *gin.Engine, db *sql.DB) {
	router.Use(middleware.ApiMiddleware(db))

	// Setup Redis session store
	token, err := handlers.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}
	store, storeError := sessions.NewRedisStore(10, "tcp", "redisDB:6379", "", []byte(token))
	if storeError != nil {
		log.Fatal("Unable to create save session with redis session: ", storeError)
	}
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		Secure:   true,
		HttpOnly: true,
	})
	router.Use(sessions.Sessions("itsmsession", store))

	// Set up middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Serve static files
	router.Static("/css", "templates/css")
	router.Static("/img", "templates/img")
	router.Static("/js", "templates/js")
	router.LoadHTMLGlob("templates/*.html")

	// Define main routes
	router.GET("/", handlers.IndexHandler)
	router.GET("/index", handlers.IndexHandler)
	router.GET("/login", handlers.LoginHandler)
	router.GET("/auth", handlers.AuthHandler)
	router.GET("/logout", handlers.LogoutHandler)
	router.GET("/login/admin", handlers.GetAgentHandler)

	// Define main router group for authorized routes
	authorized := router.Group("/")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/admin", handlers.ItsmHandler)
	}

	// Add other main groups and routes as needed
}
