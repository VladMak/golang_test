package domain

type Config struct {
	Path     string   `yaml:"path"`
	Commands []string `yaml:"commands"`
}