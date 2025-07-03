package router

import (
	"database/sql"
	"log"
	"net/http"
	"user-service/handler"
	"user-service/middleware"
	"user-service/repository"
	"user-service/services"
)

func SetupRouter(db *sql.DB) http.Handler {
	repo := &repository.UserRepo{DB: db}
	service := &services.UserService{UserRepo: repo}
	handler := &handler.UserHandler{UserService: service}

	err := repo.InitUserTable()
	if err != nil {
		log.Fatalf("DB init failed: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.Login)
	mux.HandleFunc("/protected", middleware.JWTAuth(handler.ProtectedResource))


	return mux

}
