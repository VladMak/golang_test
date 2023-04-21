package adapter

import(
	"flag"
	"fmt"
	"os"
)

type CliAdapter struct {}

// Инициализация пути до файла конфигурации
func (ca *CliAdapter) InitConfig() string {
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