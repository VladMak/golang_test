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
	"github.com/VladMak/golang_test/internal/usecase"
)

type config struct {
	Path     string   `yaml:"path"`
	Commands []string `yaml:"commands"`
}

func main() {
	cuc := usecase.ConfigUseCase{}
	cuc.CreateConfig()
	cuc.InitConfig()
	cuc.CheckConfigFile()

	// Получим аргументы из файла конфигурации
	path, commands := cuc.GetArgsFromConfigFile()

	// Начнем следить за файлами в директории
	workWithDir(path, commands)
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