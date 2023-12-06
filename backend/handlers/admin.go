package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func getAgentID(e string, c *gin.Context) (int, string, error) {
	var id int
	email := e
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, "", fmt.Errorf("unable to reach db")
	}

	err := db.QueryRow("SELECT id FROM agent WHERE email = ?", email).
		Scan(&id)
	if err != nil {
		//_, id, _ = CreateUser(c)
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent email not found in our records"})
		return 0, "", err
	}
	return id, email, nil
}

func init() {
	cid := "946670882701-dcidm9tcfdpcikpbjj8rfsb6uci22o4s.apps.googleusercontent.com"
	cs := "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"

	conf = &oauth2.Config{
		ClientID:     cid,
		ClientSecret: cs,
		RedirectURL:  "https://intel.shuttlers.africa/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

// IndexHandler handles the login
func AdminIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{})
}

// AuthHandler handles authentication of a user and initiates a session.
func AdminAuthHandler(c *gin.Context) {
	//Declare shuttlers domain
	//shuttlersDomain := "shuttlers.ng"

	// Handle the exchange code to initiate a transport.
	adminSession := sessions.Default(c)
	retrievedState := adminSession.Get("state")
	queryState := c.Request.URL.Query().Get("state")
	if retrievedState != queryState {
		log.Printf("Invalid session state: retrieved: %s; Param: %s", retrievedState, queryState)
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"message": "Invalid session state."})
		return
	}
	code := c.Request.URL.Query().Get("code")
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Login failed. Please try again."})
		return
	}

	client := conf.Client(oauth2.NoContext, tok)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	u := structs.User{}
	if err = json.Unmarshal(data, &u); err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Error marshalling response. Please try agian."})
		return
	}

	// Get user ID
	adminSession.Set("user-email", u.Email)
	adminSession.Set("user-name", u.Name)
	adminSession.Set("user-firstName", u.GivenName)
	adminSession.Set("user-lastName", u.FamilyName)
	adminSession.Set("user-sub", u.Sub)

	err = adminSession.Save()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Error while saving session. Please try again."})
		return
	}

	//
	userId, _, _ := getAgentID(u.Email, c)
	adminSession.Set("id", userId)
	adminSession.Set("user-email", u.Email)
	adminSession.Set("user-name", u.Name)
	adminSession.Set("user-firstName", u.GivenName)
	adminSession.Set("user-lastName", u.FamilyName)
	adminSession.Set("user-sub", u.Sub)
	err = adminSession.Save()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Error while saving session. Please try again."})
		return
	}
	//seen := false

	userEmail := adminSession.Get("user-Email")
	fmt.Println(userEmail)
	fmt.Println(adminSession)
	//uName := session.Get("user-name")

	c.HTML(http.StatusOK, "itsm.html", gin.H{"Username": userEmail})
	//c.HTML(http.StatusOK, "home.html", gin.H{"name": uNam, "Username": userID, "seen": seen})

}

// LoginHandler handles the login procedure.
func AgentLoginHandler(c *gin.Context) {
	state, err := RandToken(32)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Error while generating random data."})
		return
	}
	adminSession := sessions.Default(c)
	adminSession.Set("state", state)
	err = adminSession.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Error while saving session."})
		return
	}
	link := getLoginURL(state)
	c.HTML(http.StatusOK, "auth.html", gin.H{"link": link})
}

// Logout Handler
func AgentLogoutHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	adminSession.Clear()
	adminSession.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Signed out successfully",
	})
}

// RequestHandler is a rudementary handler for logged in users.
func AgentRequestHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	userEmail := adminSession.Get("user-email")
	c.HTML(http.StatusOK, "datarequest.html", gin.H{"Username": userEmail})
}

// ITSM Home
func AgentItsmHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	userEmail := adminSession.Get("user-email")
	c.HTML(http.StatusOK, "itsm.html", gin.H{"Username": userEmail})
}

// ITSM Desk
func AgentItDeskPortalHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	userEmail := adminSession.Get("user-email")
	c.HTML(http.StatusOK, "itdesk.html", gin.H{"Username": userEmail})
}
func AgentItDeskAdminHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	userEmail := adminSession.Get("user-email")
	c.HTML(http.StatusOK, "itdeskadmin.html", gin.H{"Username": userEmail})
}
func AgentItDeskHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	userEmail := adminSession.Get("user-email")
	c.HTML(http.StatusOK, "itdesk.html", gin.H{"Username": userEmail})
}

// Assets
func AgentAssetsPortalHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	userEmail := adminSession.Get("user-email")
	c.HTML(http.StatusOK, "assetsportal.html", gin.H{"Username": userEmail})
}

func AgentAssetsAdminHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	userEmail := adminSession.Get("user-email")
	c.HTML(http.StatusOK, "assetsadmin.html", gin.H{"Username": userEmail})
}
func AgentAssetsHandler(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := session.Get("user-email")
	c.HTML(http.StatusOK, "assetsx.html", gin.H{"Username": userEmail})
}

// Procurement
func ProcurementPortalHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	userEmail := adminSession.Get("user-email")
	c.HTML(http.StatusOK, "procurementportal.html", gin.H{"Username": userEmail})
}

func ProcurementAdminHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	userEmail := adminSession.Get("user-email")
	c.HTML(http.StatusOK, "procurementadmin.html", gin.H{"Username": userEmail})
}
func ProcurementHandler(c *gin.Context) {
	adminSession := sessions.Default(c)
	userEmail := adminSession.Get("user-email")
	c.HTML(http.StatusOK, "procurementx.html", gin.H{"Username": userEmail})
}
