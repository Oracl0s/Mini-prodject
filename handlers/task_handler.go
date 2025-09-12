package handlers

import (
	"encoding/json"
	"https://github.com/Oracl0s/Mini-prodject/models"
	"https://github.com/Oracl0s/Mini-prodject/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // Популярный роутер. Установите его: `go get -u github.com/gorilla/mux`
)

// TaskHandler содержит зависимости, нужные обработчикам (наше хранилище).
type TaskHandler struct {
	storage storage.Storage
}

// NewTaskHandler - конструктор для TaskHandler.
func NewTaskHandler(storage storage.Storage) *TaskHandler {
	return &TaskHandler{storage: storage}
}

// GetTasks обрабатывает GET /tasks
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := h.storage.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Кодируем список задач в JSON и отправляем клиенту
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask обрабатывает POST /tasks
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateTaskRequest

	// Парсим JSON из тела запроса в структуру CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// Передаем данные в хранилище
	newTask, err := h.storage.CreateTask(req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем созданную задачу со статусом 201 Created
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// TODO: Реализовать обработчики для GetTaskByID, UpdateTask, DeleteTask.
