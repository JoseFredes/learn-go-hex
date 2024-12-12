package main

import (
	"fmt"
	"net/http"

	sql "database/sql"

	"github.com/JoseFredes/learn-go-hex/internal/infra/left"
	"github.com/JoseFredes/learn-go-hex/internal/infra/right"
	services "github.com/JoseFredes/learn-go-hex/internal/services/user"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DNS = "user:password@tcp(127.0.0.1:3306)/db"
)

func main() {
	db, err := sql.Open("mysql", DNS)
	if err != nil {
		return
	}

	userRepository, err := right.NewMySQLUserRepository(db)
	if err != nil {
		return
	}

	userService := services.NewUserServices(userRepository)

	userHandler := left.NewUserHandler(userService)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUser(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}