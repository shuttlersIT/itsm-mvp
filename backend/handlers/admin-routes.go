package handlers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// ITSM Admin Home
func ItsmAdminHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "admin/index.html", gin.H{"Username": userEmail})
}

// ITDESK Admin Home
func ItDeskAdminHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "itdesk/admin/itdeskadmin.html", gin.H{"Username": userEmail})
}

// ITASSETS Admin Home
func AssetsAdminHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "asset/assetsadmin.html", gin.H{"Username": userEmail})
}
