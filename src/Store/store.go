package Store

import (
	"ZakirAvrora/Lab4/src/Entity"
	"ZakirAvrora/Lab4/src/e"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type Store struct {
	storagePath string
	mu          sync.Mutex
}

func New(path string) *Store {
	return &Store{storagePath: path}
}

func (s *Store) GetAllBooks() (books []Entity.Book, err error) {
	defer func() { err = e.WrapIfErr("can't get all books from store:", err) }()

	s.mu.Lock()
	defer s.mu.Unlock()
	jsonFile, err := os.Open(s.storagePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open json store: %w", err)
	}
	defer func() { err = jsonFile.Close() }()

	content, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read the contnet of json store: %w", err)
	}

	if err = json.Unmarshal(content, &books); err != nil {
		return nil, fmt.Errorf("cannot unmarshal json file: %w", err)
	}

	return books, nil
}

func (s *Store) GetBook(id int) (book Entity.Book, err error) {
	defer func() { err = e.WrapIfErr("can't get book from store:", err) }()

	books, err := s.GetAllBooks()
	if err != nil {
		return book, err
	}
	if id > len(books) || id <= 0 {
		return book, fmt.Errorf("there no book with such id")
	}
	return books[id-1], nil
}

func (s *Store) SaveBook(book Entity.Book) (newBook Entity.Book, err error) {
	defer func() { err = e.WrapIfErr("can't save book to store:", err) }()

	books, err := s.GetAllBooks()
	if err != nil {
		return newBook, err
	}
	bookId := len(books) + 1

	book.Id = bookId
	books = append(books, book)

	err = s.WriteToFile(books)
	if err != nil {
		return newBook, err
	}
	newBook = book
	return newBook, nil
}

func (s *Store) DeleteBook(id int) (err error) {
	defer func() { err = e.WrapIfErr("can't delete book from store:", err) }()

	books, err := s.GetAllBooks()
	if err != nil {
		return err
	}

	if id > len(books) || id <= 0 {
		return fmt.Errorf("id is out of range")
	}

	books = remove(books, id-1)
	err = s.WriteToFile(books)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateBook(id int, book Entity.Book) (err error) {
	defer func() { err = e.WrapIfErr("can't update book from store:", err) }()

	books, err := s.GetAllBooks()
	if err != nil {
		return err
	}

	if id > len(books) {
		return fmt.Errorf("id is out of range")
	}

	book.Id = id
	books[id-1] = book
	err = s.WriteToFile(books)
	if err != nil {
		return err
	}

	return nil

}

func remove(slice []Entity.Book, s int) []Entity.Book {
	var newSlice []Entity.Book
	removed := false

	for i := 0; i < len(slice); i++ {
		if i == s {
			removed = true
			continue
		}

		if removed {
			slice[i].Id--
		}
		newSlice = append(newSlice, slice[i])
	}

	return newSlice
}

func (s *Store) WriteToFile(books []Entity.Book) error {
	content, err := json.Marshal(books)
	if err != nil {
		return err
	}

	s.mu.Lock()
	err = os.WriteFile(s.storagePath, content, 0644)
	s.mu.Unlock()

	if err != nil {
		return err
	}
	return nil
}
