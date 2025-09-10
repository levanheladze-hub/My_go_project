package main

import (
	"fmt"
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

}

// // Выдача книги пользователю
// func (lib *Library) IssueBook(bookID, userID int) error {

// }

// // Возврат книги
// func (lib *Library) ReturnBook(bookID int) error {

// }

// // Поиск книг по автору
// func (lib *Library) FindBooksByAuthor(author string) []Book {

// }

// // Получение статистики
// func (lib *Library) GetStats() (totalBooks, issuedBooks, totalUsers int) {

// }