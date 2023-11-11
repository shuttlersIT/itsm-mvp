package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
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

	rows, err := db.Query("SELECT id, asset_id, asset_type, asset_name, description, manufacturer, model, serial_number, purchase_date, purchase_price, vendor, site, status, created_by FROM assets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var assets []structs.Asset
	for rows.Next() {
		var t structs.Asset
		if err := rows.Scan(&t.ID, &t.AssetID, &t.AssetType, &t.AssetName, &t.Description, &t.Manufacturer, &t.Model, &t.SerialNumber, &t.PurchaseDate, &t.PurchasePrice, &t.Vendor, &t.Site, &t.Status, &t.CreatedBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		assets = append(assets, t)
	}

	c.JSON(http.StatusOK, assets)
}

// Create a new asset
func CreateAsset(c *gin.Context, aid int, atype string, aname string, desc string, man string, mod string, snum string, purch ItTime, price int, ven string, sit string, statu string) {
	session := sessions.Default(c)

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
	site := sit
	status := statu
	agent := session.Get("agent-id")

	var t structs.Asset
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO assets (asset_id, asset_type, asset_name, description, manufacturer, model, serial_number, purchase_date, purchase_price, vendor, site, status, created_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", assetID, assetType, assetName, description, manufacturer, model, serialNumber, purchaseDate, purchasePrice, vendor, site, status, agent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	lastInsertID, _ := result.LastInsertId()
	t.ID = int(lastInsertID)
	c.JSON(http.StatusCreated, t)
}

// Get a asset by ID
func GetAsset(c *gin.Context, aid int) {
	id := aid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	var t structs.Asset
	err := db.QueryRow("SELECT id, asset_id, asset_type, asset_name, description, manufacturer, model, serial_number, purchase_date, purchase_price, vendor, site, status, created_by FROM assets WHERE id = ?", id).
		Scan(&t.ID, &t.AssetID, &t.AssetType, &t.AssetName, &t.Description, &t.Manufacturer, &t.Model, &t.SerialNumber, &t.PurchaseDate, &t.PurchasePrice, &t.Vendor, &t.Site, &t.Status, &t.CreatedBy)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}
	c.JSON(http.StatusOK, t)
}

// Get a asset by ID
func GetAsset2(c *gin.Context, aid int) (int, structs.Asset) {
	id := aid
	var t structs.Asset

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return 0, t
	}

	err := db.QueryRow("SELECT id, asset_id, asset_type, asset_name, description, manufacturer, model, serial_number, purchase_date, purchase_price, vendor, site, status, created_by FROM assets WHERE id = ?", id).
		Scan(&t.ID, &t.AssetID, &t.AssetType, &t.AssetName, &t.Description, &t.Manufacturer, &t.Model, &t.SerialNumber, &t.PurchaseDate, &t.PurchasePrice, &t.Vendor, &t.Site, &t.Status, &t.CreatedBy)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return 0, t
	}
	return 1, t
}

// Update a asset by ID
func UpdateAsset(c *gin.Context, i int, aid int, atype string, aname string, desc string, man string, mod string, snum string, purch ItTime, price int, ven string, sit string, statu string) {
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
	site := sit
	status := statu

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}

	var t structs.Asset
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("UPDATE assets SET asset_id = ?, asset_type = ?, asset_name = ?, description = ?, manufacturer = ?, model = ?, serial_number = ?, purchase_date = ?, purchase_price = ?, vendor = ?, site = ?, status = ? WHERE id = ?", assetID, assetType, assetName, description, manufacturer, model, serialNumber, purchaseDate, purchasePrice, vendor, site, status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Asset updated successfully")
}

// Delete a Asset by ID
func DeleteAsset(c *gin.Context, aid int) {
	id := aid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return
	}
	_, err := db.Exec("DELETE FROM assets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Asset deleted successfully")
}
