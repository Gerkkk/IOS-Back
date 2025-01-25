package main

import (
	"fmt"
	"github.com/Gerkkk/IOS-Back/main"
)

func main() {
	srv := new(Server)
	if err := srv.Run("8000", InitRoutes()); err != nil {
		fmt.Println("failed to run server")
	}
}
