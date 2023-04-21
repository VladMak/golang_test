package usecase

import (
	"github.com/VladMak/golang_test/internal/domain"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type ConfigUseCase struct {
	DomainConfig domain.Config
	filePath string
	ConfigDb domain.ConfigDb
}

func (cuc *ConfigUseCase) CreateConfig() {
	cuc.DomainConfig = domain.Config{}
	cuc.ConfigDb = domain.ConfigDb{}
}

func (cuc *ConfigUseCase) SetConfigPath(path string) {
	cuc.filePath = path
}

// Получаем аргументы из файла конфигурации
// Возможно перенести логику в драйвер, а сюда передать только параметры - TODO
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

func (cuc *ConfigUseCase) GetConfigDb() domain.ConfigDb {
	yamlFile, err := ioutil.ReadFile(cuc.filePath)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	config := cuc.ConfigDb
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Printf("DB CONFIG: %v\n", config.Db)
	return config
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