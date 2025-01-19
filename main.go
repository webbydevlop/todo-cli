package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task представляет собой структуру задачи
type Task struct {
	ID    int
	Title string
}

// Tasks — это срез (список) задач
var Tasks []Task

func main() {
	fmt.Println("Добро пожаловать в To-Do List!")
	for {
		printMenu()
		handleInput()
	}
}

// printMenu выводит меню на экран
func printMenu() {
	fmt.Println("\nВыберите действие:")
	fmt.Println("1. Показать задачи")
	fmt.Println("2. Добавить задачу")
	fmt.Println("3. Удалить задачу")
	fmt.Println("4. Выйти")
}

// handleInput обрабатывает ввод пользователя
func handleInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Убираем лишние пробелы и символы новой строки

	choice, err := strconv.Atoi(input) // Преобразуем ввод в число
	if err != nil {
		fmt.Println("Неверный выбор. Попробуйте снова.")
		return
	}

	switch choice {
	case 1:
		showTasks()
	case 2:
		addTask()
	case 3:
		fmt.Println("Удаление задачи...") // Пока заглушка
	case 4:
		fmt.Println("Выход...")
		os.Exit(0)
	default:
		fmt.Println("Неверный выбор. Попробуйте снова.")
	}
}

// showTasks выводит список задач
func showTasks() {
	if len(Tasks) == 0 {
		fmt.Println("Задач нет.")
		return
	}

	for _, task := range Tasks {
		fmt.Printf("[%d] %s\n", task.ID, task.Title)
	}
}

// addTask добавляет новую задачу
func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите название задачи: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	task := Task{
		ID:    len(Tasks) + 1, // ID задачи = количество задач + 1
		Title: title,
	}

	Tasks = append(Tasks, task)
	fmt.Println("Задача добавлена!")
}
