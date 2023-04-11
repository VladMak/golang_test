package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
	"os/exec"
)

type config struct {
	Path     string   `yaml:"path"`
	Commands []string `yaml:"commands"`
}

func main() {
	// Инициализация входных параметров (у нас это только путь до файла конфигурации)
	configFile := initConfig()

	// Проверим что файл конфига существует
	checkConfigFile(configFile)

	// Получим аргументы из файла конфигурации
	path, commands := getArgsFromConfigFile(configFile)

	// Начнем следить за файлами в директории
	workWithDir(path, commands)
}

func initConfig() string {
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
	return configFile
}

func getArgsFromConfigFile(configFile string) (string, []string) {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	config := config{}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Printf("Path: %s\nCommands: %s\n", config.Path, config.Commands)
	return config.Path, config.Commands
}

func checkConfigFile(filePath string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("Файл не найден по указанному пути: %s\n", filePath)
	} else {
		fmt.Printf("Найден конфигурационный файл: %s\n", filePath)
	}

}

func workWithDir(dirPath string, commands []string) {
	lastModified := time.Now()
	for range time.Tick(5 * time.Second) {
		// Используем Walk для прохода по всем файлам и папкам в папке.
		err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			// Пропускаем не измененные файлы и папки.
			if info.ModTime().Before(lastModified) {
				return nil
			}

			// Выводим сообщение об изменениях.
			fmt.Printf("%s was modified\n", path)
			// Необходимо запустить команды из YAML файла
			runCmd(commands)

			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		lastModified = time.Now()
	}

}

func runCmd(commands []string) {
	for _, cmd := range commands {
		execCmd := exec.Command("sh", "-c", cmd)
		err := execCmd.Run()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}
	fmt.Println("Commands Done!")
}