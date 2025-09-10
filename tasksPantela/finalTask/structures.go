package main

type Book struct {
    ID        int
    Title     string
    Author    string
    Year      int
    IsIssued  bool
}

type User struct {
    ID    int
    Name  string
    Email string
}

type Library struct {
    Books      []Book
    Users      []User
    Issues     map[int]int // bookID -> userID
}

type LibraryManager interface {
    AddBook(book Book) error
    RegisterUser(user User) error
    IssueBook(bookID, userID int) error
    ReturnBook(bookID int) error
    FindBooksByAuthor(author string) []Book
    GetStats() (int, int, int)
}