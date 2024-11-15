package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JoseFredes/learn-go-hex/internal/infra/left"
	"github.com/JoseFredes/learn-go-hex/internal/infra/right"
	services "github.com/JoseFredes/learn-go-hex/internal/services/user"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/mydatabase" 
	
	userRepository, err := right.NewMySQLUserRepository(dsn)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}


	userService := services.NewUserServices(userRepository)


	userHandler := left.NewUserHandler(userService)

	http.HandleFunc("/users/", userHandler.GetUser)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}