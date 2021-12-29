package service

import (
	"selling-management-be/helper/pagination"
	"selling-management-be/model"
	"strings"
)

type OrganizationListRequest struct {
	Ids             []string `json:"ids"`
	Status          string   `json:"status"`
	Keyword         string   `json:"keyword"`
	pagination.Pagination
}

func OrganizationList(request *OrganizationListRequest) (reply []*OrganizationGetReply, err error) {
	sql := mainService.db
	if len(request.Ids) > 0 {
		sql = sql.Where("id IN (?)", request.Ids)
	}

	if len(request.Ids) > 0 {
		sql = sql.Where("organization_id IN (?)", request.Ids)
	}

	if len(request.Status) > 0 {
		sql = sql.Where("status = ?", request.Status)
	}

	if len(request.Keyword) > 0 {
		lowerKey := strings.ToLower(request.Keyword)
		sql = sql.Where("(LOWER(name) ILIKE ?)", "%"+lowerKey+"%")
	}

	sql = sql.Offset(request.Limit * request.Page).
		Limit(request.Limit).
		Order("id DESC")

	var organizations []*model.Organization
	if err = sql.Find(&organizations).Error; err != nil {
		return nil, err
	}

	reply, err = toOrganizationListReply(organizations)
	return
}

func toOrganizationListReply(organizations []*model.Organization) ([]*OrganizationGetReply, error) {
	var organizationGetReplies []*OrganizationGetReply
	for _, organization := range organizations {
		organizationGetReply, err := toOrganizationGetReply(organization)
		if err != nil {
			return nil, err
		}
		organizationGetReplies = append(organizationGetReplies, organizationGetReply)
	}
	return organizationGetReplies, nil
}

