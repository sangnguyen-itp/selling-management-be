package service

import (
	"github.com/mitchellh/mapstructure"
	"selling-management-be/model"
)

type OrganizationGetRequest struct {
	ID string `json:"id"`
}

type OrganizationGetReply struct {
	ID      string `gorm:"primaryKey"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func OrganizationGet(request *OrganizationGetRequest) (reply *OrganizationGetReply, err error) {
	var organization model.Organization

	if err = mainService.db.First(&organization, "id = ?", request.ID).Error; err != nil {
		return nil, err
	}
	reply, err = toOrganizationGetReply(&organization)
	return
}

func toOrganizationGetReply(organization *model.Organization) (*OrganizationGetReply, error) {
	var reply OrganizationGetReply
	if err := mapstructure.Decode(organization, &reply); err != nil {
		return nil, err
	}
	return &reply, nil
}
