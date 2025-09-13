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