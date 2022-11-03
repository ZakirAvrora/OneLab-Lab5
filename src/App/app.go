package App

import (
	"ZakirAvrora/OneLab-lab5/src/Entity"
	"ZakirAvrora/OneLab-lab5/src/Repository"
	"ZakirAvrora/OneLab-lab5/src/Store"
	"errors"
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
		if errors.Is(err, Store.ErrNoRowAffected) {
			return c.String(http.StatusBadRequest, "Bad request, no affect on data")
		}
		return err
	}
	return c.String(http.StatusAccepted, "Successfully updated")
}

func (app *App) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "the id parameter must be integer")
	}
	if err = app.repo.DeleteBook(id); err != nil {
		if errors.Is(err, Store.ErrNoRowAffected) {
			return c.String(http.StatusBadRequest, "Bad request, no affect on data")
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusAccepted, "Successfully deleted")
}
