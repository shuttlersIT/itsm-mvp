package routers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/middleware"
)

// SetupAPIRouter sets up the API router
func SetupAPIRouter(api *gin.Engine, db *sql.DB) {
	api.Use(middleware.ApiMiddleware(db))

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
	api.Use(sessions.Sessions("itsmsession", store))

	// Set up middleware
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	// Define API routes
	api.GET("/", handlers.IndexHandler)
	api.GET("/index", handlers.IndexHandler)
	api.GET("/login", handlers.LoginHandler)
	api.GET("/auth", handlers.AuthHandler)
	api.GET("/logout", handlers.LogoutHandler)
	api.GET("/login/admin", handlers.GetAgentHandler)

	// Define API router group for authorized routes
	authorized := api.Group("/itsm")
	authorized.Use(middleware.AuthorizeRequest())
	{
		authorized.GET("/itsm/tickets", handlers.ItsmHandler)
		authorized.GET("/itsm/assets", handlers.ItsmHandler)
		authorized.GET("/itsm/procurement", handlers.ItsmHandler)
		authorized.GET("/itsm/admin", handlers.ItsmHandler)
	}

	// Add other API groups and routes as needed
}
