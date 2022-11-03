package App

import (
	"ZakirAvrora/Lab4/src/Entity"
	"ZakirAvrora/Lab4/src/Repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type App struct {
	repo Repository.Repository
}

func New(repository Repository.Repository) *App {
	return &App{repo: repository}
}

func (app *App) GetBooks(c echo.Context) error {
	books, err := app.repo.GetAllBooks()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSONPretty(http.StatusOK, books, " ")
}

func (app *App) GetBookByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "the id parameter must be integer")
	}
	book, err := app.repo.GetBook(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSONPretty(http.StatusOK, book, " ")
}

func (app *App) SaveBook(c echo.Context) error {
	book := &Entity.Book{}
	if err := c.Bind(book); err != nil {
		return err
	}

	newBook, err := app.repo.SaveBook(*book)
	if err != nil {
		return err
	}

	return c.JSONPretty(http.StatusCreated, newBook, " ")
}

func (app *App) UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "the id parameter must be integer")
	}

	book := &Entity.Book{}
	if err := c.Bind(book); err != nil {
		return err
	}

	if err := app.repo.UpdateBook(id, *book); err != nil {
		return err
	}
	return c.String(http.StatusNoContent, "Successfully updated")
}

func (app *App) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "the id parameter must be integer")
	}
	if err = app.repo.DeleteBook(id); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusNoContent, "Successfully deleted")
}
