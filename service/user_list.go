package service

import (
	"selling-management-be/helper"
	"selling-management-be/model"
)

type UserListRequest struct {
	Ids       []string `json:"ids"`
	Usernames []string `json:"usernames"`
	Status    string   `json:"status"`
	Keyword   string   `json:"keyword"`
	helper.Pagination
}

func UserList(request *UserListRequest) (reply []*UserGetReply, err error) {
	sql := mainService.db
	if len(request.Ids) > 0 {
		sql = sql.Where("id IN (?)", request.Ids)
	}

	if len(request.Usernames) > 0 {
		sql = sql.Where("username IN (?)", request.Usernames)
	}

	if len(request.Status) > 0 {
		sql = sql.Where("status = ?", request.Usernames)
	}

	if len(request.Keyword) > 0 {
		sql = sql.Where("email ILIKE ?", "%" + request.Keyword + "%")
	}

	sql = sql.Offset(request.Limit * request.Page).
			Limit(request.Limit).
			Order("id DESC")

	var users []*model.User
	if err = sql.Find(&users).Error; err != nil {
		return nil, err
	}

	reply, err = toUserListReply(users)
	return
}

func toUserListReply(users []*model.User) ([]*UserGetReply, error) {
	var userGetReplies []*UserGetReply
	for _, user := range users {
		userGetReply, err := toUserGetReply(user)
		if err != nil {
			return nil, err
		}
		userGetReplies = append(userGetReplies, userGetReply)
	}
	return userGetReplies, nil
}