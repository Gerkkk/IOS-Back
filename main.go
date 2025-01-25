package main

import (
	"fmt"
	gotype "github.com/Gadzet005/GoType/backend"
)

// @title GoType App API
// @version 0.0.1
// @description API Server for GoType game and website

// @host localhost:8000
// @BasePath/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	srv := new(main.Server)
	if err := srv.Run(8000, handlers.InitRoutes()); err != nil {
		fmt.Println("failed to run server")
	}
}
