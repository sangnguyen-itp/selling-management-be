package service

import (
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"selling-management-be/model"
	"strings"
	"time"
)

type ProductUpdateRequest struct {
	ID             string          `json:"id"`
	Code           string          `json:"code"`
	Type           string          `json:"type"`
	Name           string          `json:"name"`
	WholesalePrice decimal.Decimal `json:"wholesale_price"`
	RetailPrice    decimal.Decimal `json:"retail_price"`
	RetailUnit     string          `json:"retail_unit"`
	WholesaleUnit  string          `json:"wholesale_unit"`
	Currency       string          `json:"currency"`
	Status         string          `json:"status"`
	OrganizationID string          `json:"-"`
	UpdatedBy      string          `json:"-"`
	UpdatedAt      time.Time       `json:"-"`
}

type ProductUpdateReply struct {
	ID string `json:"id"`
}

func ProductUpdate(request *ProductUpdateRequest) (reply *ProductUpdateReply, err error) {
	var product model.Product
	if err = mapstructure.Decode(request, &product); err != nil {
		return nil, err
	}

	sqlDB := mainService.db.Model(&product)
	if len(request.OrganizationID) > 0 {
		sqlDB = sqlDB.Where("organization_id = ?", request.OrganizationID)
	}

	productUpdateData := getProductUpdateDataMap(request)
	if err = sqlDB.Select(allowZero()).Updates(productUpdateData).Error; err != nil {
		return nil, err
	}

	return &ProductUpdateReply{ID: product.ID}, nil
}

func allowZero() string {
	var allowCols []string
	allowCols = append(allowCols, "price")
	return strings.Join(allowCols, ",")
}

func getProductUpdateDataMap(request *ProductUpdateRequest) map[string]interface{} {
	data := make(map[string]interface{})
	if len(request.Code) > 0 {
		data["code"] = request.Code
	}

	if len(request.Name) > 0 {
		data["name"] = request.Name
	}

	if !request.RetailPrice.IsZero() {
		data["retail_price"] = request.RetailPrice
	}

	if !request.RetailPrice.IsZero() {
		data["wholesale_price"] = request.WholesalePrice
	}

	if len(request.RetailUnit) > 0 {
		data["retail_unit"] = request.RetailPrice
	}

	if len(request.RetailUnit) > 0 {
		data["wholesale_unit"] = request.WholesaleUnit
	}

	if len(request.Currency) > 0 {
		data["currency"] = request.Currency
	}

	if len(request.Status) > 0 {
		data["status"] = request.Status
	}

	return data
}
