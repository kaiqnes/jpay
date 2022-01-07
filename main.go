package main

import (
	"fmt"
	"github.com/nuno/nunes-jumia/src/app"
	"os"
)

func main() {
	router := app.SetupApp()

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
