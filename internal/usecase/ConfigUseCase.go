package usecase

import (
	"github.com/VladMak/golang_test/internal/domain"
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type ConfigUseCase struct {
	DomainConfig domain.Config
	filePath string
}

func (cuc *ConfigUseCase) CreateConfig() {
	cuc.DomainConfig = domain.Config{}
}


// Инициализация пути до файла конфигурации
func (cuc *ConfigUseCase) InitConfig() {
	// Определим переменные для параметров.
	var configFile string

	// Добавим параметры в флаги.
	flag.StringVar(&configFile, "c", "./config/config.yaml", "Config file")

	// Разобрать аргументы.
	flag.Parse()

	// Выведем параметры.
	fmt.Println("Config file:", configFile)

	// Проверим аргументы.
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}
	cuc.filePath = configFile
}

// Получаем аргументы из файла конфигурации
func (cuc *ConfigUseCase) GetArgsFromConfigFile() (string, []string) {
	yamlFile, err := ioutil.ReadFile(cuc.filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	config := cuc.DomainConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Printf("Path: %s\nCommands: %s\n", config.Path, config.Commands)
	return config.Path, config.Commands
}

// Проверка на существование файла конфигурации
// Если файл есть - true, нет - false - TODO
func (cuc *ConfigUseCase) CheckConfigFile() {
	_, err := os.Stat(cuc.filePath)
	if os.IsNotExist(err) {
		fmt.Printf("Файл не найден по указанному пути: %s\n", cuc.filePath)
	} else {
		fmt.Printf("Найден конфигурационный файл: %s\n", cuc.filePath)
	}

}