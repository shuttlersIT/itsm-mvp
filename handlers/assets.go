package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/structs"
)

// List all Assets
func ListAssets(c *gin.Context) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	rows, err := db.Query("SELECT id, asset_id, asset_type, asset_name, description, manufacturer, model, serial_number, purchase_date, purchase_price, vendor, site, status, description, status FROM assets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var assets []structs.Asset
	for rows.Next() {
		var t structs.Asset
		if err := rows.Scan(&t.ID, &t.Subject, &t.Description, &t.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tickets = append(assets, t)
	}

	c.JSON(http.StatusOK, assets)
}

// Create a new ticket
func CreateAsset(c *gin.Context, aid int, atype string, aname string, desc string, man string, mod string, snum string, purch handlers.ItTime, price float, ven string, site string, status string) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	assetID := aid      
	assetType := atype
	assetName := aname
	description := desc
	manufacturer := man
	model := mod
	serialNumber := snum
	purchaseDate := purch
	purchasePrice := price
	vendor := ven
	site := site
	status := status

	var t structs.Ticket
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO tickets (title, description, status) VALUES (?, ?, ?)", t.Subject, t.Description, t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	lastInsertID, _ := result.LastInsertId()
	t.ID = int(lastInsertID)
	c.JSON(http.StatusCreated, t)
}

// Get a ticket by ID
func GetAsset(c *gin.Context) {
	id := c.Param("id")

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	var t structs.Ticket
	err := db.QueryRow("SELECT id, subject, description, status FROM tickets WHERE id = ?", id).
		Scan(&t.ID, &t.Subject, &t.Description, &t.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	c.JSON(http.StatusOK, t)
}

// Update a ticket by ID
func UpdateAsset(c *gin.Context, i int, aid int, atype string, aname string, desc string, man string, mod string, snum string, purch handlers.ItTime, price float, ven string, site string, status string) {
	id := i
	assetID := aid      
	assetType := atype
	assetName := aname
	description := desc
	manufacturer := man
	model := mod
	serialNumber := snum
	purchaseDate := purch
	purchasePrice := price
	vendor := ven
	site := site
	status := status

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	var t structs.Ticket
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE tickets SET title = ?, description = ?, status = ? WHERE id = ?", t.Subject, t.Description, t.Status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Ticket updated successfully")
}

// Delete a ticket by ID
func DeleteAsset(c *gin.Context) {
	id := c.Param("id")

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	_, err := db.Exec("DELETE FROM tickets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Ticket deleted successfully")
}
