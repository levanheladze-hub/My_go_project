package main

import (
	// "fmt"
	"log"
)

func main() {
    library := &Library{
        Books:  []Book{},
        Users:  []User{},
        Issues: make(map[int]int),
    }

    // Добавляем книги
    books := []Book{
        {ID: 1, Title: "1984", Author: "George Orwell", Year: 1949},
        {ID: 2, Title: "Brave New World", Author: "Aldous Huxley", Year: 1932},
        {ID: 3, Title: "Как вырастить Зураба в пробирке", Author: "Хеладзе Леван", Year: 2025},
		{ID: 4, Title: "Как вырастить Зураба в пробирке", Author: "Хеладзе Леван", Year: 2025},
		{ID: 5, Title: "Как отпиздить корейца", Author: "Хеладзе Леван", Year: 2024}, 
		{ID: 5, Title: "Как отпиздить корейца", Author: "Хеладзе Леван", Year: 2024}, 
    }
    
    for _, book := range books {
        if err := library.AddBook(book); err != nil {
            log.Printf("Ошибка добавления книги: %v\n", err)
        }
    }

    // Регистрируем пользователей
    users := []User{
        {ID: 1, Name: "Иван Иванов", Email: "ivan@mail.com"},
        {ID: 2, Name: "Петр Петров", Email: "petr@mail.com"},
        {ID: 3, Name: "Шин Александр", Email: "kon-chen-i@mail.com"},
		{ID: 4, Name: "Учадзе Зураб", Email: "konchita@mail.com"},
    }
    
    for _, user := range users {
        if err := library.RegisterUser(user); err != nil {
            log.Printf("Ошибка регистрации: %v\n", err)
        }
    }

    // Выдаем книги
    if err := library.IssueBook(1, 1); err != nil {
        log.Printf("Ошибка выдачи: %v\n", err)
    }

    if err := library.IssueBook(2, 2); err != nil {
        log.Printf("Ошибка выдачи: %v\n", err)
    }

    // Пытаемся выдать уже выданную книгу
    if err := library.IssueBook(1, 2); err != nil {
        log.Printf("Ошибка: %v\n", err) // Должна быть ошибка
    }

    // Ищем книги Orwell
    levanBooks := library.FindBooksByAuthor("Хеладзе Леван")
    log.Println("Книги Хеладзе Левана:", len(levanBooks))
    for _, book := range levanBooks {
        log.Println("Книги автора Хеладзе Левана: ", book.Title)
    }

    // Получаем статистику
    total, issued, usersCount := library.GetStats()
    log.Printf("Всего книг: %d, Выдано: %d, Пользователей: %d\n", total, issued, usersCount)

    // Возвращаем книгу
    if err := library.ReturnBook(1); err != nil {
        log.Printf("Ошибка возврата: %v\n", err)
    }

    // Проверяем статистику после возврата
    _, issued, _ = library.GetStats()
    log.Printf("После возврата - Выдано: %d\n", issued)

    // Выдаем книги
    if err := library.IssueBook(1, 3); err != nil {
        log.Printf("Ошибка выдачи: %v\n", err)
    }

    // Получаем статистику
    total, issued, usersCount = library.GetStats()
    log.Printf("Всего книг: %d, Выдано: %d, Пользователей: %d\n", total, issued, usersCount)
}