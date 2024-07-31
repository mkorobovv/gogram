package insta_provider

import (
	"github.com/Davincible/goinsta/v3"
	"github.com/mkorobovv/gogram/internal/app/domain"
)

func toEntity(dto *goinsta.Users) []domain.InstagramUser {
	users := make([]domain.InstagramUser, len(dto.Users))

	for i, user := range dto.Users {
		users[i] = instaUserToEntity(user)
	}

	return users
}

func instaUserToEntity(dto *goinsta.User) domain.InstagramUser {
	return domain.InstagramUser{
		ID:          dto.ID,
		Username:    dto.Username,
		FullName:    dto.FullName,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
	}
}
