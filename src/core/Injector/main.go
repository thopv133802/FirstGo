package Injector

import (
	"GoFirst/src/controllers"
	"GoFirst/src/domain/repositories"
	"GoFirst/src/domain/services"
	"github.com/jinzhu/gorm"
)

func NewCommonController() (controllers.CommonControllerImpl, error) {
	service, err := NewCommonService()
	return controllers.CommonControllerImpl{
		Service: service,
	}, err
}

func NewCommonService() (services.CommonServiceImpl, error) {
	repository, err := NewCommonRepository("", "")
	return services.CommonServiceImpl{
		Repository: repository,
	}, err
}

func NewCommonRepository(databaseName, connectionString string) (*repositories.CommonRepositoryImpl, error) {
	database, err := gorm.Open(databaseName, connectionString)
	return &repositories.CommonRepositoryImpl{
		DB: database,
	}, err
}