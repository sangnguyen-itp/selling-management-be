package resource

import (
	"selling-management-be/conf"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/user_status"
	"selling-management-be/pkg/logger"
	"selling-management-be/service"
	"time"
)

func initSystemData() {
	id := "user-121ee30d-972c-477b-ac58-91b4efa6c479"
	organizationID := "*"
	_, err := service.UserCreate(&service.UserCreateRequest{
		ID:             id,
		FirstName:      conf.EnvConfig.DefaultFirstName,
		LastName:       conf.EnvConfig.DefaultLastName,
		Email:          conf.EnvConfig.DefaultEmail,
		PhoneNumber:    conf.EnvConfig.DefaultPhoneNumber,
		Address:        conf.EnvConfig.DefaultAddress,
		Username:       conf.EnvConfig.DefaultUsername,
		Password:       conf.EnvConfig.DefaultPassword,
		Role:           conf.EnvConfig.DefaultRole,
		OrganizationID: organizationID,
		LoginTime:      time.Now(),
		Status:         user_status.Active,
		IsSystem:       true,
		UpdatedBy:      "system",
		CreatedBy:      "system",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	})
	if err != nil {
		logger.Log().Error(domain.UserDomain, "service.UserCreate", err)
		return
	}
}
