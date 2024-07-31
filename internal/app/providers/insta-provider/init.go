package insta_provider

import (
	"log"

	"github.com/Davincible/goinsta/v3"
	"github.com/mkorobovv/gogram/internal/app/domain/config"
)

type InstaProvider struct {
	config         config.InstaProvider
	accountManager accountManager
}

type accountManager interface {
	Following(query string, order goinsta.FollowOrder) *goinsta.Users
	Followers(query string) *goinsta.Users
}

func New(config config.InstaProvider) *InstaProvider {
	instagram := goinsta.New(config.BasicAuth.Username, config.BasicAuth.Password)

	err := instagram.Login()
	if err != nil {
		log.Fatalf("error login: %v", err)
	}

	return &InstaProvider{
		config:         config,
		accountManager: instagram.Account,
	}
}
