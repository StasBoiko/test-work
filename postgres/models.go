// Code generated by sqlc. DO NOT EDIT.

package postgres

import ()

type Author struct {
	ID   int64
	Name string
}

type Book struct {
	ID    int64
	Title string
}

type BookAuthor struct {
	ID       int64
	BookID   int64
	AuthorID int64
}
