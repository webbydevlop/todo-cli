package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Task представляет собой структуру задачи
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Tasks — это срез (список) задач
var Tasks []Task

func main() {
	fmt.Println("Добро пожаловать в To-Do List!")

	// Загружаем задачи из файла при запуске
	loadTasks()

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
	fmt.Println("4. Отметить задачу как выполненную")
	fmt.Println("5. Поиск задач по названию")
	fmt.Println("6. Выйти")
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
		deleteTask()
	case 4:
		completeTask()
	case 5:
		searchTasks()
	case 6:
		fmt.Println("Выход...")
		saveTasks() // Сохраняем задачи перед выходом
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
		status := " "
		if task.Completed {
			status = "✓"
		}
		fmt.Printf("[%d] %s %s\n", task.ID, status, task.Title)
	}
}

// addTask добавляет новую задачу
func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите название задачи: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	task := Task{
		ID:        len(Tasks) + 1, // ID задачи = количество задач + 1
		Title:     title,
		Completed: false, // По умолчанию задача не выполнена
	}

	Tasks = append(Tasks, task)
	fmt.Println("Задача добавлена!")
}

// deleteTask удаляет задачу по ID
func deleteTask() {
	if len(Tasks) == 0 {
		fmt.Println("Задач нет.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите ID задачи для удаления: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input) // Преобразуем ввод в число
	if err != nil {
		fmt.Println("Неверный ID. Попробуйте снова.")
		return
	}

	// Ищем задачу по ID
	for i, task := range Tasks {
		if task.ID == id {
			// Удаляем задачу из среза
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			fmt.Println("Задача удалена!")
			return
		}
	}

	fmt.Println("Задача с таким ID не найдена.")
}

// completeTask отмечает задачу как выполненную
func completeTask() {
	if len(Tasks) == 0 {
		fmt.Println("Задач нет.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите ID задачи для отметки как выполненной: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input) // Преобразуем ввод в число
	if err != nil {
		fmt.Println("Неверный ID. Попробуйте снова.")
		return
	}

	// Ищем задачу по ID
	for i, task := range Tasks {
		if task.ID == id {
			// Отмечаем задачу как выполненную
			Tasks[i].Completed = true
			fmt.Println("Задача отмечена как выполненная!")
			return
		}
	}

	fmt.Println("Задача с таким ID не найдена.")
}

// searchTasks ищет задачи по ключевому слову
func searchTasks() {
	if len(Tasks) == 0 {
		fmt.Println("Задач нет.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите ключевое слово для поиска: ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)

	found := false
	for _, task := range Tasks {
		if strings.Contains(strings.ToLower(task.Title), strings.ToLower(keyword)) {
			status := " "
			if task.Completed {
				status = "✓"
			}
			fmt.Printf("[%d] %s %s\n", task.ID, status, task.Title)
			found = true
		}
	}

	if !found {
		fmt.Println("Задачи не найдены.")
	}
}

// saveTasks сохраняет задачи в файл
func saveTasks() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Ошибка при сохранении задач:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Добавляем форматирование для читаемости
	if err := encoder.Encode(Tasks); err != nil {
		fmt.Println("Ошибка при сохранении задач:", err)
	}
}

// loadTasks загружает задачи из файла
func loadTasks() {
	file, err := os.Open("tasks.json")
	if err != nil {
		// Если файл не существует, просто возвращаемся
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Tasks); err != nil {
		fmt.Println("Ошибка при загрузке задач:", err)
	}
}
