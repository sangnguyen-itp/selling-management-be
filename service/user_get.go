package service

import "selling-management-be/model"

type UserGetRequest struct {
	ID         string
}

func GetUser(request *UserGetRequest) (reply *model.User, err error) {
	if err = mainService.db.First(reply, "id = ?", request.ID).Error; err != nil {
		return nil, err
	}

	return reply, nil
}
