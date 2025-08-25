package config

type ClientConfig struct {
	YouTube struct {
		ID          string `yaml:"id"`
		Secret      string `yaml:"secret"`
		RedirectURL string `yaml:"redirect_url"`
	} `yaml:"youtube"`
}

type Config struct {
	Client ClientConfig `yaml:"client"`
}
