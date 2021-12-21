package service

import (
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"selling-management-be/model"
)

type ProductGetRequest struct {
	ID string `json:"id"`
}

type ProductGetReply struct {
	ID       string          `gorm:"primaryKey"`
	Code     string          `json:"code"`
	Name     string          `json:"name"`
	Price    decimal.Decimal `json:"price"`
	Currency string          `json:"currency"`
}

func ProductGet(request *ProductGetRequest) (reply *ProductGetReply, err error) {
	var product model.Product
	if err = mainService.db.First(&product, "id = ?", request.ID).Error; err != nil {
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

	reply.Price = product.Price
	return &reply, nil
}
