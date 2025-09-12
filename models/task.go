package models

// Импортируем пакет для работы со временем
import "time"

// Task - структура, описывающая задачу
type Task struct {
	ID        int       `json:"id"`         // Уникальный идентификатор
	Title     string    `json:"title"`      // Название задачи
	Completed bool      `json:"completed"`  // Статус выполнения
	CreatedAt time.Time `json:"created_at"` // Время создания
}

// Данные для создания новой задачи (тело POST запроса)
type CreateTaskRequest struct {
	Title string `json:"title"`
}
