package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nuno/nunes-jumia/src/controller"
	"github.com/nuno/nunes-jumia/src/repository"
	"github.com/nuno/nunes-jumia/src/service"
	"gorm.io/gorm"
)

func SetupResources(dbSession *gorm.DB) *gin.Engine {
	customerRepository := repository.NewCustomerRepository(dbSession)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	return setupRoutes(customerController)
}

func setupRoutes(controller controller.CustomerController) *gin.Engine {
	router := gin.New()

	router.GET("/customer/search", controller.GetCustomers)

	return router
}
