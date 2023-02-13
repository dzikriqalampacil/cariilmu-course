package helper

import (
	"github.com/dzikriqalampacil/cariilmu-course/model/domain"
	"github.com/dzikriqalampacil/cariilmu-course/model/web"
)

func ToUserResponse(User domain.User) web.UserResponse {
	return web.UserResponse{
		Id:    User.Id,
		Name:  User.Name,
		Email: User.Email,
	}
}

func ToUserResponses(categories []domain.User) []web.UserResponse {
	var UserResponses []web.UserResponse
	for _, User := range categories {
		UserResponses = append(UserResponses, ToUserResponse(User))
	}
	return UserResponses
}
