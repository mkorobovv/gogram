package service

import "github.com/mkorobovv/gogram/internal/app/domain"

type Service struct {
	instaProvider instaProvider
}

type instaProvider interface {
	GetFollowings() []domain.InstagramUser
	GetFollowers() []domain.InstagramUser
}

func New(instaProvider instaProvider) *Service {
	return &Service{
		instaProvider: instaProvider,
	}
}
