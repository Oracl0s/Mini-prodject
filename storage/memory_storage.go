package storage

import (
	"https://github.com/Oracl0s/Mini-prodject/models"
	"time"
) // Замените "your-module-name" на имя вашего модуля

// Storage - интерфейс, описывающий все методы для работы с данными.
// Это позволит в будущем легко заменить хранилище.
type Storage interface {
	GetAllTasks() ([]models.Task, error)
	CreateTask(title string) (*models.Task, error)
	// ... другие методы (GetByID, Update, Delete)
}

// MemoryStorage - реализация Storage, хранящая данные в памяти.
type MemoryStorage struct {
	tasks  []models.Task
	lastID int // Счетчик для присвоения ID
}

// NewMemoryStorage - конструктор, создающий новое хранилище в памяти.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks:  make([]models.Task, 0),
		lastID: 0,
	}
}

// GetAllTasks возвращает все задачи.
func (s *MemoryStorage) GetAllTasks() ([]models.Task, error) {
	return s.tasks, nil
}

// CreateTask создает новую задачу и добавляет ее в хранилище.
func (s *MemoryStorage) CreateTask(title string) (*models.Task, error) {
	s.lastID++
	task := models.Task{
		ID:        s.lastID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	s.tasks = append(s.tasks, task)
	return &task, nil
}

// TODO: Реализовать остальные методы (GetTaskByID, UpdateTask, DeleteTask)
