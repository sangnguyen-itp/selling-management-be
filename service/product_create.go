package service

import (
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/user_status"
	"selling-management-be/model"
	"selling-management-be/pkg/generate_id"
	"time"
)

type ProductCreateRequest struct {
	Name           string          `json:"name"`
	Code           string          `json:"code"`
	Price          decimal.Decimal `json:"price"`
	Currency       string          `json:"currency"`
	Status         string          `json:"status"`
	OrganizationID string          `json:"organization_id"`
	UpdatedBy      string          `json:"-"`
	UpdatedAt      time.Time       `json:"-"`
	CreatedBy      string          `json:"-"`
	CreatedAt      time.Time       `json:"-"`
}

type ProductCreateReply struct {
	ID string `json:"id"`
}

func ProductCreate(request *ProductCreateRequest) (reply *ProductCreateReply, err error) {
	var product model.Product
	if err = mapstructure.Decode(request, &product); err != nil {
		return nil, err
	}

	id := generate_id.NewID(domain.ProductDomain)
	product.ID = id
	product.SearchName = product.Name
	product.Price = request.Price
	product.Status = user_status.Active

	if err = mainService.db.Create(&product).Error; err != nil {
		return nil, err
	}

	return &ProductCreateReply{ID: product.ID}, nil
}
