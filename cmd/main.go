package main

import (
	"github.com/Oracl0s/Mini-prodject/handlers"
	"github.com/Oracl0s/Mini-prodject/storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
import "fmt"

func main() {
	// Создаем роутер
	r := mux.NewRouter()

	// Инициализируем наше "хранилище" в памяти
	storage := storage.NewMemoryStorage()
	// Инициализируем обработчики, передав им хранилище
	taskHandler := handlers.NewTaskHandler(storage)

	// Регистрируем HTTP маршруты
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	// r.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	// r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	// r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	// Запускаем сервер на порту 8080
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Hello World")
}
