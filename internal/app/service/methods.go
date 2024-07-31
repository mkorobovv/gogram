package service

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/mkorobovv/gogram/internal/app/domain"
)

func (svc *Service) Start() {
	ticker := time.NewTicker(120 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var wg sync.WaitGroup

		wg.Add(2)

		var (
			followings []domain.InstagramUser
			followers  []domain.InstagramUser
		)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			followings = svc.instaProvider.GetFollowings()
		}(&wg)

		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			followers = svc.instaProvider.GetFollowers()
		}(&wg)

		wg.Wait()

		followerMap := make(map[int64]struct{}, len(followers))

		for _, follower := range followers {
			followerMap[follower.ID] = struct{}{}
		}

		unfollowed := make([]domain.InstagramUser, 0)

		for _, following := range followings {
			if _, found := followerMap[following.ID]; !found {
				unfollowed = append(unfollowed, following)
			}
		}

		if len(unfollowed) == 0 {
			continue
		}

		log.Printf("%d unfollowed accounts!", len(unfollowed))

		err := writeUnfollowed(unfollowed)
		if err != nil {
			log.Println(err)

			continue
		}

		log.Println("accounts successfully written")
	}
}

const filename = "unfollowed.txt"

func writeUnfollowed(unfollowed []domain.InstagramUser) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755) //nolint:gosec
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	for _, unfollow := range unfollowed {
		_, err = file.WriteString(fmt.Sprintf("ID: %d, Username: %s, FullName: %s\n",
			unfollow.ID,
			unfollow.Username,
			unfollow.FullName,
		))
		if err != nil {
			return err
		}
	}

	return nil
}
