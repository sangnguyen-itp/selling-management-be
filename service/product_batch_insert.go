package service

import (
	"github.com/mitchellh/mapstructure"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/product_status"
	"selling-management-be/model"
	"selling-management-be/pkg/generate_id"
)

type ProductBatchInsertRequest struct {
	Products []*ProductCreateRequest
}

type ProductBatchInsertReply struct {}

func ProductBatchInsert(request *ProductBatchInsertRequest) (*ProductBatchInsertReply, error){
	if request != nil && len(request.Products) > 0 {
		tx := mainService.db.Begin()
		var products []*model.Product
		for _, p := range request.Products {
			var product model.Product
			if err := mapstructure.Decode(p, &product); err != nil {
				return nil, err
			}

			id := generate_id.NewID(domain.ProductDomain)
			product.ID = id
			product.SearchName = product.Name
			product.Price = p.Price
			product.Status = product_status.Active

			products = append(products, &product)
		}

		if err := tx.Create(products).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		tx.Commit()
	}

	return &ProductBatchInsertReply{}, nil
}
