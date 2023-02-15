package main

import (
    "go-app-users"
    "log"
    "go-app-users/pkg/handler"
)

func main() {
    handlers := new(handler.Handler)
    srv := new(users.Server)
    if err := srv.Run("9000", handlers.InitRoutes()); err != nil {
        log.Fatalf("Error occured while running http server: %s", err.Error())
    }
}
