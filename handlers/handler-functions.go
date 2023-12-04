package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/database"
	"github.com/shuttlersIT/itsm-mvp/stafftokenization"
	"github.com/shuttlersIT/itsm-mvp/structs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func TokenTimestamp() int64 {
	return time.Now().Unix()
}

// var c string
var conf *oauth2.Config

// RandToken generates a random @l length token.
func RandToken(l int) (string, error) {
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// Share databases
func ShareDb(d *sql.DB, c *gin.Context) *sql.DB {
	db := d
	database.TableExists(db, "tickets")
	return d
}

func getID(e string, c *gin.Context) (int, string, error) {
	var id int
	email := e
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, "", fmt.Errorf("cant reach db")
	}

	err := db.QueryRow("SELECT id FROM staff WHERE email = ?", email).
		Scan(&id)
	if err != nil {
		//_, id, _ = CreateUser2(c)
		c.JSON(http.StatusNotFound, gin.H{"error": "Staff email not found in our records"})
		return 0, "", err
	}
	return id, email, nil
}

func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
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
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// AuthHandler handles authentication of a user and initiates a session.
func AuthHandler(c *gin.Context) {
	//Declare shuttlers domain
	//shuttlersDomain := "shuttlers.ng"

	// Handle the exchange code to initiate a transport.
	userSession := sessions.Default(c)
	userToken := userSession.Get("user-token-gen-on-server-side")

	if userToken != nil {
		// Retrieve user_id from session
		token, _ := userToken.(string)
		userName, _ := userSession.Get("user-email").(string)

		_, er := stafftokenization.ParseToken(token)
		if er != nil {
			log.Printf("token doesnt exist")
		}

		if len(userName) > 0 {
			log.Printf("token already exists")
			// Run the LogoutHandler first
			LogoutHandler(c)
		}

	}
	retrievedState := userSession.Get("state")
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

	//
	userId, _, _ := getID(u.Email, c)
	//seen := false
	userEmail := userSession.Get("user-Email")
	_, token, e := JwtToken(userId, u.Email)
	if e != nil {
		log.Fatal("unable to set token in session", e)
	}

	userSession.Set("user-name", u.Name)
	userSession.Set("user-token-gen-on-server-side", token)
	userSession.Set("id", userId)
	userSession.Set("user-email", u.Email)
	userSession.Set("user-firstName", u.GivenName)
	userSession.Set("user-lastName", u.FamilyName)
	userSession.Set("user-sub", u.Sub)
	err = userSession.Save()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"message": "Error while saving session. Please try again."})
		return
	}

	// Return a response with the generated token
	c.JSON(http.StatusOK, gin.H{"access_token": token, "user_email": u.Email})
	c.HTML(http.StatusOK, "itsm.html", gin.H{"Username": userEmail})
	//c.HTML(http.StatusOK, "home.html", gin.H{"name": uNam, "Username": userID, "seen": seen})

}

// LoginHandler handles the login procedure.
func LoginHandler(c *gin.Context) {
	startState, err := RandToken(64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Error while generating random data."})
		return
	}
	userSession := sessions.Default(c)
	userSession.Set("state", startState)
	err = userSession.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Error while saving session."})
		return
	}
	link := getLoginURL(startState)
	c.HTML(http.StatusOK, "auth.html", gin.H{"link": link})
}

// Logout Handler
func LogoutHandler(c *gin.Context) {
	userSession := sessions.Default(c)
	userSession.Clear()
	userSession.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Signed out successfully",
	})
}

// RequestHandler is a rudementary handler for logged in users.
func RequestHandler(c *gin.Context) {
	userSession := sessions.Default(c)
	userEmail := userSession.Get("user-email")
	c.HTML(http.StatusOK, "datarequest.html", gin.H{"Username": userEmail})
}

// ITSM Home
func ItsmPortalHandler(c *gin.Context) {
	userSession := sessions.Default(c)
	userEmail := userSession.Get("user-email")
	c.HTML(http.StatusOK, "index.html", gin.H{"Username": userEmail})
}

func ItsmHandler(c *gin.Context) {
	userSession := sessions.Default(c)
	userEmail := userSession.Get("user-email")
	c.HTML(http.StatusOK, "asset/assetsx.html", gin.H{"Username": userEmail})
}

// ITSM Desk
func ItDeskPortalHandler(c *gin.Context) {
	userSession := sessions.Default(c)
	userEmail := userSession.Get("user-email")
	c.HTML(http.StatusOK, "itdesk/itdesk.html", gin.H{"Username": userEmail})
}

func ItDeskHandler(c *gin.Context) {
	userSession := sessions.Default(c)
	userEmail := userSession.Get("user-email")
	c.HTML(http.StatusOK, "itdesk/itdesk.html", gin.H{"Username": userEmail})
}

// Assets
func AssetsPortalHandler(c *gin.Context) {
	userSession := sessions.Default(c)
	userEmail := userSession.Get("user-email")
	c.HTML(http.StatusOK, "asset/assetsportal.html", gin.H{"Username": userEmail})
}

func AssetsHandler(c *gin.Context) {
	userSession := sessions.Default(c)
	userEmail := userSession.Get("user-email")
	c.HTML(http.StatusOK, "asset/assetsx.html", gin.H{"Username": userEmail})
}
