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

type ProductBatchInsertReply struct{}

func ProductBatchPut(request *ProductBatchInsertRequest) (*ProductBatchInsertReply, error) {
	if request != nil && len(request.Products) > 0 {
		tx := mainService.db.Begin()
		var products []*model.Product
		for _, p := range request.Products {
			product, err := isUpdateProduct(p)
			if err != nil {
				return nil, err
			}
			products = append(products, product)
		}

		if len(products) > 0 {
			if err := tx.Save(&products).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
		}

		tx.Commit()
	}

	return &ProductBatchInsertReply{}, nil
}

func isUpdateProduct(p *ProductCreateRequest) (*model.Product, error) {
	var putProduct model.Product
	if err := mapstructure.Decode(p, &putProduct); err != nil {
		return nil, err
	}

	putProducts, err := ProductList(&ProductListRequest{
		Codes: []string{p.Code},
	})
	if err != nil {
		return nil, err
	}

	if len(putProducts) > 0 {
		if err = mapstructure.Decode(putProducts[0], &putProduct); err != nil {
			return nil, err
		}

		putProduct.RetailPrice = p.RetailPrice
		putProduct.WholesalePrice = p.WholesalePrice
		return &putProduct, nil
	}

	id := generate_id.NewID(domain.ProductDomain)
	putProduct.ID = id
	putProduct.SearchName = p.Name
	putProduct.RetailPrice = p.RetailPrice
	putProduct.WholesalePrice = p.WholesalePrice
	putProduct.Status = product_status.Active

	return &putProduct, nil
}
