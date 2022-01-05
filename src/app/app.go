package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nuno/nunes-jumia/src/config"
	"github.com/nuno/nunes-jumia/src/controller"
	"github.com/nuno/nunes-jumia/src/repository"
	"github.com/nuno/nunes-jumia/src/service"
	"gorm.io/gorm"
)

func SetupApp() *gin.Engine {
	dbSession := config.GetDatabase()

	return setupResources(dbSession)
}

func setupResources(dbSession *gorm.DB) *gin.Engine {
	router := gin.Default()

	customerRepository := repository.NewCustomerRepository(dbSession)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	return customerController.SetupRoutes(router)
}
