package insta_provider

import (
	"github.com/Davincible/goinsta/v3"
	"github.com/mkorobovv/gogram/internal/app/domain"
)

const defaultInstagramLimit = 200

func (prv *InstaProvider) GetFollowings() []domain.InstagramUser {
	users := make([]domain.InstagramUser, 0, defaultInstagramLimit)

	userFollowings := prv.accountManager.Following("", goinsta.LatestOrder)

	prevNextID := ""

	for userFollowings.Next() {
		users = append(users, toEntity(userFollowings)...)

		if prevNextID == userFollowings.NextID {
			break
		}

		prevNextID = userFollowings.NextID
	}

	return users
}

func (prv *InstaProvider) GetFollowers() []domain.InstagramUser {
	users := make([]domain.InstagramUser, 0, defaultInstagramLimit)

	userFollowers := prv.accountManager.Followers("")

	prevNextID := ""

	for userFollowers.Next() {
		users = append(users, toEntity(userFollowers)...)

		if prevNextID == userFollowers.NextID {
			break
		}

		prevNextID = userFollowers.NextID
	}

	return users
}
