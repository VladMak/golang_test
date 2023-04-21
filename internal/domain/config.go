package domain

type Config struct {
	Path     string   `yaml:"path"`
	Commands []string `yaml:"commands"`
}

type ConfigDb struct {
	Db struct {
		Username string `yaml:"username"`
		Password string	`yaml:"password"`
		Host     string `yaml:"host"`
		Port     int `yaml:"port"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"db"`
}