package main

import "fmt"

// Объявляем интерфейс. Он определяет поведение: "что-то, что умеет говорить".
type Speaker interface {
	Speak() string // Любой тип, у которого есть метод Speak() string, удовлетворяет интерфейсу Speaker.
}

// Тип Dog
type Dog struct{ Name string }

// Тип Dog реализует метод Speak, значит, он удовлетворяет интерфейсу Speaker.
func (d Dog) Speak() string {
	return "Гав! Меня зовут " + d.Name
}

// Тип Cat
type Cat struct{ Name string }

// Тип Cat тоже реализует метод Speak, значит, он тоже удовлетворяет интерфейсу Speaker.
func (c Cat) Speak() string {
	return "Мяу! Я " + c.Name
}

// Функция, принимающая не конкретный тип, а интерфейс.
// Она может работать с ЛЮБЫМ типом, который умеет "Speak".
func Introduce(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{Name: "Шарик"}
	cat := Cat{Name: "Мурка"}

	// Мы можем передать и Dog, и Cat в функцию Introduce, т.к. оба они реализуют Speaker.
	Introduce(dog) // Вывод: Гав! Меня зовут Шарик
	Introduce(cat) // Вывод: Мяу! Я Мурка

	// Можно создать слайс интерфейса, чтобы хранить в нем разные типы.
	zoo := []Speaker{dog, cat}
	for _, animal := range zoo {
		fmt.Println(animal.Speak())
	}
}
