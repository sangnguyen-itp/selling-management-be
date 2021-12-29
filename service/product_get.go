package service

import (
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"selling-management-be/model"
)

type ProductGetRequest struct {
	ID             string `json:"id"`
	OrganizationID string `json:"-"`
}

type ProductGetReply struct {
	ID             string          `gorm:"primaryKey"`
	Code           string          `json:"code" gorm:"index:idx_code,unique"`
	Type           string          `json:"type"`
	Name           string          `json:"name"`
	SearchName     string          `json:"search_name" gorm:"index:idx_search_name"`
	WholesalePrice decimal.Decimal `json:"wholesale_price" gorm:"type:decimal(20,0);"`
	RetailPrice    decimal.Decimal `json:"retail_price" gorm:"type:decimal(20,0);"`
	RetailUnit     string          `json:"retail_unit"`
	WholesaleUnit  string          `json:"wholesale_unit"`
	Currency string          `json:"currency"`
	Status         string          `json:"status"`
	OrganizationID string          `json:"organization_id" gorm:"index:idx_product_organization"`
}

func ProductGet(request *ProductGetRequest) (reply *ProductGetReply, err error) {
	var product model.Product

	sqlDB := mainService.db
	if len(request.OrganizationID) > 0 {
		sqlDB = sqlDB.Where("organization_id = ?", request.OrganizationID)
	}

	if err = sqlDB.First(&product, "id = ?", request.ID).Error; err != nil {
		return nil, err
	}
	reply, err = toProductGetReply(&product)
	return
}

func toProductGetReply(product *model.Product) (*ProductGetReply, error) {
	var reply ProductGetReply
	if err := mapstructure.Decode(product, &reply); err != nil {
		return nil, err
	}

	reply.RetailPrice = product.RetailPrice
	reply.WholesalePrice = product.WholesalePrice
	return &reply, nil
}
