package main

import (
	"fmt"
	"github.com/nuno/nunes-jumia/src/config"
	"github.com/nuno/nunes-jumia/src/routes"
	"os"
)

func main() {
	dbSession := config.GetDatabase()

	router := routes.SetupResources(dbSession)

	port := getPort()
	if err := router.Run(":" + port); err != nil {
		fmt.Printf("error to start the server: %v\n", err)
	}
}

func getPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return
}
