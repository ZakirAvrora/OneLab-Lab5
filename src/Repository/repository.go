package Repository

import "ZakirAvrora/OneLab-lab5/src/Entity"

type Repository interface {
	GetAllBooks() ([]Entity.Book, error)
	GetBook(id int) (Entity.Book, error)
	SaveBook(book Entity.Book) (Entity.Book, error)
	UpdateBook(id int, book Entity.Book) error
	DeleteBook(id int) error
}
