package config

type Config struct {
	Application Application `yaml:"application"`
	Providers   Providers   `yaml:"providers"`
}

type Application struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Providers struct {
	InstaProvider InstaProvider `yaml:"instaProvider"`
}

type InstaProvider struct {
	BasicAuth BasicAuth `yaml:"basicAuth"`
}

type BasicAuth struct {
	Username string `env:"INSTA_PROVIDER_USERNAME" env-required:"true" yaml:"username"`
	Password string `env:"INSTA_PROVIDER_PASSWORD" env-required:"true" yaml:"password"`
}
