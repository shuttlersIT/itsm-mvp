package converter

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shuttlersIT/itsm-mvp/handlers"
	"github.com/shuttlersIT/itsm-mvp/structs"
	frontendstructs "github.com/shuttlersIT/itsm-mvp/structs/frontend"
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
