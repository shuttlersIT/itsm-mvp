package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
)

// List all Assets
func ListAssets(c *gin.Context) ([]*structs.Asset, bool, error) {
	status := false
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, status, fmt.Errorf("unable to reach db")
	}

	rows, err := db.Query("SELECT id, asset_id, asset_type, asset_name, description, manufacturer, model, serial_number, purchase_date, purchase_price, vendor, site, status, created_by FROM assets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, status, fmt.Errorf("unable to find asset")
	}
	defer rows.Close()

	var assets []*structs.Asset
	for rows.Next() {
		var t *structs.Asset
		if err := rows.Scan(&t.ID, &t.AssetID, &t.AssetType, &t.AssetName, &t.Description, &t.Manufacturer, &t.Model, &t.SerialNumber, &t.PurchaseDate, &t.PurchasePrice, &t.Vendor, &t.Site, &t.Status, &t.CreatedBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil, status, fmt.Errorf("unable to retrieve asset")
		}
		assets = append(assets, t)
	}

	if len(assets) < 1 {
		status = true
	}

	c.JSON(http.StatusOK, assets)
	return assets, status, nil
}

// Create a new asset
func CreateAsset(c *gin.Context, a structs.Asset) (*structs.Asset, int, error) {
	session := sessions.Default(c)

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, 0, fmt.Errorf("db unreacheable")
	}
	//assetID := a.ID
	assetnum := a.AssetID
	assetType := a.AssetType
	assetName := a.AssetName
	description := a.Description
	manufacturer := a.Manufacturer
	model := a.Model
	serialNumber := a.SerialNumber
	purchaseDate := a.PurchaseDate
	purchasePrice := a.PurchasePrice
	vendor := a.Vendor
	site := a.Site
	status := a.Status
	agent := session.Get("agent-id")

	var t structs.Asset
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, 0, fmt.Errorf("asset failed json bind")
	}

	result, err := db.Exec("INSERT INTO assets (asset_id, asset_type, asset_name, description, manufacturer, model, serial_number, purchase_date, purchase_price, vendor, site, status, created_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", assetnum, assetType, assetName, description, manufacturer, model, serialNumber, purchaseDate, purchasePrice, vendor, site, status, agent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, 0, fmt.Errorf("asset Creation failed")
	}

	lastInsertID, _ := result.LastInsertId()
	t.ID = int(lastInsertID)
	c.JSON(http.StatusCreated, t)

	c.JSON(http.StatusOK, "User created successfully")
	return &t, t.ID, nil
}

// Get a asset by ID
func GetAsset(c *gin.Context, aid int) (*structs.Asset, error) {
	id := aid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach db")
	}
	var t structs.Asset
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("asset failed json bind")
	}

	err := db.QueryRow("SELECT id, asset_id, asset_type, asset_name, description, manufacturer, model, serial_number, purchase_date, purchase_price, vendor, site, status, created_by FROM assets WHERE id = ?", id).
		Scan(&t.ID, &t.AssetID, &t.AssetType, &t.AssetName, &t.Description, &t.Manufacturer, &t.Model, &t.SerialNumber, &t.PurchaseDate, &t.PurchasePrice, &t.Vendor, &t.Site, &t.Status, &t.CreatedBy)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return nil, fmt.Errorf("unable to find asset")
	}
	c.JSON(http.StatusOK, t)

	return &t, nil
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
func UpdateAsset(c *gin.Context, a structs.Asset) (*structs.Asset, error) {
	ID := a.ID
	assetID := a.AssetID
	assetType := a.AssetType
	assetName := a.AssetName
	description := a.Description
	manufacturer := a.Manufacturer
	model := a.Model
	serialNumber := a.SerialNumber
	purchaseDate := a.PurchaseDate
	purchasePrice := a.PurchasePrice
	vendor := a.Vendor
	site := a.Site
	status := a.Status

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return nil, fmt.Errorf("unable to reach db")
	}

	var t structs.Asset
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("unable to bind json")
	}

	result, err := db.Exec("UPDATE assets SET asset_id = ?, asset_type = ?, asset_name = ?, description = ?, manufacturer = ?, model = ?, serial_number = ?, purchase_date = ?, purchase_price = ?, vendor = ?, site = ?, status = ? WHERE id = ?", assetID, assetType, assetName, description, manufacturer, model, serialNumber, purchaseDate, purchasePrice, vendor, site, status, ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, fmt.Errorf("unable to update asset")
	}

	lastInsertID, _ := result.LastInsertId()
	newAssetID := int(lastInsertID)
	newAsset, _ := GetAsset(c, newAssetID)

	c.JSON(http.StatusOK, "Asset updated successfully")
	return newAsset, nil
}

// Delete a Asset by ID
func DeleteAsset(c *gin.Context, aid int) error {
	id := aid

	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get user handler"})
		return fmt.Errorf("unable to reach db")
	}
	_, err := db.Exec("DELETE FROM assets WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return fmt.Errorf("unable to delete ticket")
	}
	c.JSON(http.StatusOK, "Asset deleted successfully")
	return nil
}

func updateAssetType(c *gin.Context, t structs.AssetType) (string, error) {
	// Don't forget type assertion when getting the connection from context.
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unable to reach DB from get update user handler"})
		return "failed", fmt.Errorf("can't reach db")
	}

	t.AssetTypeID, _ = strconv.Atoi(c.PostForm("assetTypeID"))
	t.AssetType = c.PostForm("typeName")

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return "failed", fmt.Errorf("unable to bind json")
	}

	_, err := db.Exec("UPDATE asset_type SET asset_type = ? WHERE id = ?", t.AssetType, t.AssetTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "failed", fmt.Errorf("unable to update asset")
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Asset type updated successfully",
	})

	return "Asset Type Update Successful", nil
}
