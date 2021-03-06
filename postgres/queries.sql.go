// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package postgres

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, name)
	var i Author
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createBook = `-- name: CreateBook :one
INSERT INTO books (title)
VALUES ($1)
RETURNING id, title
`

func (q *Queries) CreateBook(ctx context.Context, title string) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook, title)
	var i Book
	err := row.Scan(&i.ID, &i.Title)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name FROM authors
ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAuthorsByBookID = `-- name: ListAuthorsByBookID :many
SELECT authors.id, authors.name FROM authors, book_authors
WHERE authors.id = book_authors.author_id AND book_authors.book_id = $1
`

func (q *Queries) ListAuthorsByBookID(ctx context.Context, bookID int64) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthorsByBookID, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBooks = `-- name: ListBooks :many
SELECT id, title FROM books
ORDER BY title
`

func (q *Queries) ListBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBooksByAuthorID = `-- name: ListBooksByAuthorID :many
SELECT books.id, books.title FROM books, book_authors
WHERE books.id = book_authors.book_id AND book_authors.author_id = $1
`

func (q *Queries) ListBooksByAuthorID(ctx context.Context, authorID int64) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooksByAuthorID, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(&i.ID, &i.Title); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setBookAuthor = `-- name: SetBookAuthor :exec
INSERT INTO book_authors (book_id, author_id)
VALUES ($1, $2)
`

type SetBookAuthorParams struct {
	BookID   int64
	AuthorID int64
}

func (q *Queries) SetBookAuthor(ctx context.Context, arg SetBookAuthorParams) error {
	_, err := q.db.ExecContext(ctx, setBookAuthor, arg.BookID, arg.AuthorID)
	return err
}

const unsetBookAuthors = `-- name: UnsetBookAuthors :exec
DELETE FROM book_authors
WHERE book_id = $1
`

func (q *Queries) UnsetBookAuthors(ctx context.Context, bookID int64) error {
	_, err := q.db.ExecContext(ctx, unsetBookAuthors, bookID)
	return err
}

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE authors
SET name = $2
WHERE id = $1
RETURNING id, name
`

type UpdateAuthorParams struct {
	ID   int64
	Name string
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, updateAuthor, arg.ID, arg.Name)
	var i Author
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const updateBook = `-- name: UpdateBook :one
UPDATE books
SET title = $2
WHERE id = $1
RETURNING id, title
`

type UpdateBookParams struct {
	ID    int64
	Title string
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, updateBook, arg.ID, arg.Title)
	var i Book
	err := row.Scan(&i.ID, &i.Title)
	return i, err
}
