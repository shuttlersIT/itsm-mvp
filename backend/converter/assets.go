package converter

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/backend/handlers"
	"github.com/shuttlersIT/itsm-mvp/backend/structs"
	frontendstructs "github.com/shuttlersIT/itsm-mvp/backend/structs/frontend"
)

func FrontEndAsset(c *gin.Context, t structs.Asset) *frontendstructs.FrontendAsset {
	var asset frontendstructs.FrontendAsset
	asset.ID = t.ID
	asset.AssetName = t.AssetName
	asset.AssetType = t.AssetType
	asset.Description = t.Description
	asset.Manufacturer = t.Manufacturer
	asset.Model = t.Model
	asset.SerialNumber = t.SerialNumber
	asset.Site = t.Site
	asset.Status = t.Status
	asset.Vendor = t.Vendor
	asset.PurchaseDate = t.PurchaseDate
	asset.PurchasePrice = t.PurchasePrice
	i, _ := handlers.GetAgent(c, t.CreatedBy)
	asset.CreatedBy = fmt.Sprintf("%v %v", i.FirstName, i.LastName)
	asset.CreatedAt = t.CreatedAt
	asset.UpdatedAt = t.UpdatedAt

	return &asset
}

func FrontEndAssetList(c *gin.Context, t structs.Asset) *frontendstructs.FrontendAsset {
	var asset frontendstructs.FrontendAsset
	asset.ID = t.ID
	asset.AssetName = t.AssetName
	asset.AssetType = t.AssetType
	asset.Description = t.Description
	asset.Manufacturer = t.Manufacturer
	asset.Model = t.Model
	asset.SerialNumber = t.SerialNumber
	asset.Site = t.Site
	asset.Status = t.Status
	asset.Vendor = t.Vendor
	asset.PurchaseDate = t.PurchaseDate
	asset.PurchasePrice = t.PurchasePrice
	i, _ := handlers.GetAgent(c, t.CreatedBy)
	asset.CreatedBy = fmt.Sprintf("%v %v", i.FirstName, i.LastName)
	asset.CreatedAt = t.CreatedAt
	asset.UpdatedAt = t.UpdatedAt

	return &asset
}

// FrontEndAsset efficiently converts a structs.Asset into a frontendstructs.FrontendAsset
func FrontEndAssetB(c *gin.Context, t *structs.Asset) *frontendstructs.FrontendAsset {
	var asset frontendstructs.FrontendAsset
	asset.ID = t.ID
	asset.AssetName = t.AssetName
	asset.AssetType = t.AssetType
	asset.Description = t.Description
	asset.Manufacturer = t.Manufacturer
	asset.Model = t.Model
	asset.SerialNumber = t.SerialNumber
	asset.Site = t.Site
	asset.Status = t.Status
	asset.Vendor = t.Vendor
	asset.PurchaseDate = t.PurchaseDate
	asset.PurchasePrice = t.PurchasePrice
	i, _ := handlers.GetAgent(c, t.CreatedBy)
	asset.CreatedBy = fmt.Sprintf("%v %v", i.FirstName, i.LastName)
	asset.CreatedAt = t.CreatedAt
	asset.UpdatedAt = t.UpdatedAt

	return &asset
}

// FrontEndAssetList efficiently converts a slice of structs.Asset into a slice of frontendstructs.FrontendAsset
func FrontEndAssetListB(c *gin.Context, assetList []*structs.Asset) []*frontendstructs.FrontendAsset {
	frontendAssetList := make([]*frontendstructs.FrontendAsset, len(assetList))

	for i, a := range assetList {
		frontendAssetList[i] = FrontEndAssetB(c, a)
	}

	return frontendAssetList
}

func ConvertFrontendToAsset(c *gin.Context, ft *frontendstructs.FrontendAsset) (*structs.Asset, error) {
	var a structs.Asset
	// Populate fields of Ticket using data from FrontendTicket
	// Fetch additional details using handler functions if needed
	// Handle errors appropriately

	a.ID = ft.ID
	a.AssetName = ft.AssetName
	a.AssetType = ft.AssetType
	a.AssetID = ft.AssetID
	a.Description = ft.Description
	a.Manufacturer = ft.Manufacturer
	a.Model = ft.Model
	a.SerialNumber = ft.SerialNumber
	a.PurchaseDate = ft.PurchaseDate
	a.PurchasePrice = ft.PurchasePrice
	a.Vendor = ft.Vendor
	a.Site = ft.Site
	a.Status = ft.Status
	//agent, _ := handlers.GetAgent(c, ft.CreatedBy)
	//a.CreatedBy = agent.AgentID
	// ... (populate other fields)

	return &a, nil
}
