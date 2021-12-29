package service

import (
	"github.com/mitchellh/mapstructure"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/organization_status"
	"selling-management-be/model"
	"selling-management-be/pkg/generate_id"
	"time"
)

type OrganizationCreateRequest struct {
	ID      string `gorm:"primaryKey"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	UpdatedBy      string          `json:"-"`
	UpdatedAt      time.Time       `json:"-"`
	CreatedBy      string          `json:"-"`
	CreatedAt      time.Time       `json:"-"`
}

type OrganizationCreateReply struct {
	ID string `json:"id"`
}

func OrganizationCreate(request *OrganizationCreateRequest) (reply *OrganizationCreateReply, err error) {
	var organization model.Organization
	if err = mapstructure.Decode(request, &organization); err != nil {
		return nil, err
	}

	id := generate_id.NewID(domain.ProductDomain)
	organization.ID = id
	organization.Status = organization_status.Active

	if err = mainService.db.Create(&organization).Error; err != nil {
		return nil, err
	}

	return &OrganizationCreateReply{ID: organization.ID}, nil
}

