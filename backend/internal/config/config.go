package config

type ClientConfig struct {
	Google struct {
		ID          string `yaml:"id"`
		Secret      string `yaml:"secret"`
		RedirectURL string `yaml:"redirect_url"`
	} `yaml:"Google"`
}

type Config struct {
	Client ClientConfig `yaml:"client"`
}
