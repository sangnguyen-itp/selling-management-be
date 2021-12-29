package model

import (
	"encoding/json"
	"selling-management-be/defined/domain"
	"selling-management-be/pkg/logger"
	"time"
)

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	ID             string          `gorm:"primaryKey"`
	Code           string          `json:"code" gorm:"index:idx_code,unique"`
	Type           string          `json:"type"`
	Name           string          `json:"name"`
	SearchName     string          `json:"search_name" gorm:"index:idx_search_name"`
	Price          decimal.Decimal `json:"price" gorm:"type:decimal(20,8);"`
	Currency       string          `json:"currency"`
	Status         string          `json:"status"`
	OrganizationID string          `json:"organization_id" gorm:"index:idx_product_organization"`

	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Encode() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		logger.Log().Error(domain.ProductDomain, "Encode failed", err)
		return ""
	}
	return string(bytes)
}

func (p *Product) Decode(bytes string) error {
	err := json.Unmarshal([]byte(bytes), p)
	logger.Log().Error(domain.ProductDomain, "Decode failed", err)
	return err
}

func (p *Product) BuildSearch() {

}
