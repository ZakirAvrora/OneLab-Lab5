package Store

import (
	"ZakirAvrora/OneLab-lab5/src/Entity"
	"ZakirAvrora/OneLab-lab5/src/e"
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

var ErrNoRowAffected = errors.New("bad request, no affect in data")

type Store struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllBooks() (books []Entity.Book, err error) {
	defer func() { err = e.WrapIfErr("can't get all books from store:", err) }()
	err = s.db.Select(&books, `SELECT * from books`)

	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *Store) GetBook(id int) (book Entity.Book, err error) {
	defer func() { err = e.WrapIfErr("can't get book from store:", err) }()

	err = s.db.Get(&book, "SELECT * FROM books WHERE id= $1", id)
	if err != nil {
		return
	}
	return book, nil
}

func (s *Store) SaveBook(book Entity.Book) (newBook Entity.Book, err error) {
	defer func() { err = e.WrapIfErr("can't save book to store:", err) }()
	query := `INSERT INTO books (title, author, price)
			  VALUES (:title, :author, :price)`
	_, err = s.db.NamedExec(query, book)
	if err != nil {
		return Entity.Book{}, err
	}
	newBook = book

	return newBook, nil
}

func (s *Store) DeleteBook(id int) (err error) {
	defer func() { err = e.WrapIfErr("can't delete book from store:", err) }()
	result, err := s.db.Exec(`DELETE FROM books WHERE id = $1`, id)

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNoRowAffected
	}
	return err
}

func (s *Store) UpdateBook(id int, book Entity.Book) (err error) {
	defer func() { err = e.WrapIfErr("can't update book from store:", err) }()

	result, err := s.db.Exec(`UPDATE books SET title = $2, author= $3, price=$4, created_at=$5
			WHERE id=$1`, id, book.Title, book.Author, book.Price, time.Now())

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNoRowAffected
	}
	return err
}
