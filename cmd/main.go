package main

import (
	insta_provider "github.com/mkorobovv/gogram/internal/app/providers/insta-provider"
	"github.com/mkorobovv/gogram/internal/app/service"
	"github.com/mkorobovv/gogram/internal/pkg/config"
)

func main() {
	cfg := config.New()

	instaProvider := insta_provider.New(cfg.Providers.InstaProvider)

	svc := service.New(instaProvider)

	svc.Start()
}
