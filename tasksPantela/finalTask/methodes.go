package main

import (
	"fmt"
    "sort"
)

// Добавление новой книги
func (lib *Library) AddBook(book Book) error {
    // Проверяем, есть ли уже книга с таким ID
    for _, existingBook := range lib.Books {
        if existingBook.ID == book.ID {
            return fmt.Errorf("книга с ID %d уже существует", book.ID)
        }
    }
    // Если ID уникален - добавляем книгу
    lib.Books = append(lib.Books, book)
    return nil
}

// Регистрация нового пользователя  
func (lib *Library) RegisterUser(user User) error {
    // Проверяем, есть ли уже книга с таким ID
    for _, existingUser := range lib.Users {
        if existingUser.ID == user.ID {
            return fmt.Errorf("читатель с ID %d уже существует", user.ID)
        }
    }
    // Если ID уникален - добавляем книгу
    lib.Users = append(lib.Users, user)
    return nil
}

// Выдача книги пользователю
func (lib *Library) IssueBook(bookID, userID int) error {
    // Сортируем книги по ID для бинарного поиска
    sort.Slice(lib.Books, func(i, j int) bool {
        return lib.Books[i].ID < lib.Books[j].ID
    })

    // Бинарный поиск книги
    left, right := 0, len(lib.Books)-1
    bookIndex := -1
    var foundBook *Book

    for left <= right {
        mid := left + (right-left)/2
        currentBook := lib.Books[mid]
        
        if currentBook.ID == bookID {
            bookIndex = mid
            foundBook = &lib.Books[mid]
            break
        }
        
        if currentBook.ID < bookID {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    if foundBook == nil {
        return fmt.Errorf("книги с ID %d не найдено", bookID)
    }

    if foundBook.IsIssued {
        return fmt.Errorf("книга с ID %d уже выдана", bookID)
    }

    // Сортируем пользователей по ID для бинарного поиска
    sort.Slice(lib.Users, func(i, j int) bool {
        return lib.Users[i].ID < lib.Users[j].ID
    })

    // Бинарный поиск пользователя
    left, right = 0, len(lib.Users)-1
    userFound := false

    for left <= right {
        mid := left + (right-left)/2
        currentUser := lib.Users[mid]
        
        if currentUser.ID == userID {
            userFound = true
            break
        }
        
        if currentUser.ID < userID {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    if !userFound {
        return fmt.Errorf("читателя с ID %d не найдено", userID)
    }

    // Выдаем книгу
    lib.Issues[bookID] = userID
    lib.Books[bookIndex].IsIssued = true
    return nil
}

// Возврат книги
func (lib *Library) ReturnBook(bookID int) error {
    delete(lib.Issues, bookID)
    for index, book := range lib.Books {
        if book.ID == bookID {
            lib.Books[index].IsIssued = false
        } else {
            continue
        }
    }
    return nil
}

// Поиск книг по автору
func (lib *Library) FindBooksByAuthor(author string) []Book {
    var Books []Book
    for _, book := range lib.Books {
        if book.Author == author {
            Books = append(Books, book)
        } 
    }
    return Books
}

// Получение статистики
func (lib *Library) GetStats() (totalBooks, issuedBooks, totalUsers int) {
    totalBooks = len(lib.Books)
    totalUsers = len(lib.Users)
    issuedBooks = len(lib.Issues)
    return totalBooks, issuedBooks, totalUsers
}