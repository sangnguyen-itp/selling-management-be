package service

import (
	"selling-management-be/helper"
	"selling-management-be/model"
	"strings"
)

type ProductListRequest struct {
	Ids             []string `json:"ids"`
	OrganizationIds []string `json:"organization_ids"`
	Status          string   `json:"status"`
	Keyword         string   `json:"keyword"`
	helper.Pagination
}

func ProductList(request *ProductListRequest) (reply []*ProductGetReply, err error) {
	sql := mainService.db
	if len(request.Ids) > 0 {
		sql = sql.Where("id IN (?)", request.Ids)
	}

	if len(request.OrganizationIds) > 0 {
		sql = sql.Where("organization_id IN (?)", request.OrganizationIds)
	}

	if len(request.Status) > 0 {
		sql = sql.Where("status = ?", request.Status)
	}

	if len(request.Keyword) > 0 {
		lowerKey := strings.ToLower(request.Keyword)
		sql = sql.Where("(LOWER(id) ILIKE ? OR LOWER(search_name) ILIKE ? OR LOWER(code) ILIKE ?)", "%"+lowerKey+"%", "%"+lowerKey+"%", "%"+lowerKey+"%")
	}

	sql = sql.Offset(request.Limit * request.Page).
		Limit(request.Limit).
		Order("id DESC")

	var products []*model.Product
	if err = sql.Find(&products).Error; err != nil {
		return nil, err
	}

	reply, err = toProductListReply(products)
	return
}

func toProductListReply(products []*model.Product) ([]*ProductGetReply, error) {
	var productGetReplies []*ProductGetReply
	for _, product := range products {
		productGetReply, err := toProductGetReply(product)
		if err != nil {
			return nil, err
		}
		productGetReplies = append(productGetReplies, productGetReply)
	}
	return productGetReplies, nil
}
