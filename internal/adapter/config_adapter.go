package adapter

import "github.com/VladMak/golang_test/internal/usecase"

type ConfigAdapter struct{
	cuc usecase.ConfigUseCase
}

func (ca *ConfigAdapter) SetConfigPath(configPath string, cuc *usecase.ConfigUseCase) {
	cuc.SetConfigPath(configPath)
}