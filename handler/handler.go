package handler

import "selling-management-be/service"

func ValidateOrganizationID(organizationID string) bool {
	_, err := service.OrganizationGet(&service.OrganizationGetRequest{
		ID: organizationID,
	})
	return err == nil
}
