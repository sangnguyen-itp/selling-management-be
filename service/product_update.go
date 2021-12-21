package service

import (
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"selling-management-be/model"
	"strings"
	"time"
)

type ProductUpdateRequest struct {
	ID        string          `json:"id"`
	Code      string          `json:"code"`
	Name      string          `json:"name"`
	Price     decimal.Decimal `json:"price"`
	Currency  string          `json:"currency"`
	Status    string          `json:"status"`
	UpdatedBy string          `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}

type ProductUpdateReply struct {
	ID string `json:"id"`
}

func ProductUpdate(request *ProductUpdateRequest) (reply *ProductUpdateReply, err error) {
	var product model.Product
	if err = mapstructure.Decode(request, &product); err != nil {
		return nil, err
	}

	productUpdateData := getProductUpdateDataMap(request)
	if err = mainService.db.Model(&product).Select(allowZero()).Updates(productUpdateData).Error; err != nil {
		return nil, err
	}

	return &ProductUpdateReply{ID: product.ID}, nil
}

func allowZero() string {
	var allowCols []string
	allowCols = append(allowCols, "price")
	return strings.Join(allowCols, ",")
}

func getProductUpdateDataMap(request *ProductUpdateRequest) (data map[string]interface{}) {
	if len(request.Code) > 0 {
		data["code"] = request.Code
	}

	if len(request.Name) > 0 {
		data["name"] = request.Name
	}

	if !request.Price.IsZero() {
		data["price"] = request.Price
	}

	if len(request.Currency) > 0 {
		data["currency"] = request.Currency
	}

	if len(request.Status) > 0 {
		data["status"] = request.Status
	}

	return
}
