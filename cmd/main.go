package main

import (
	"ZakirAvrora/OneLab-lab5/src/App"
	"ZakirAvrora/OneLab-lab5/src/Store"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
)

/*lab4:
создать template crud используя фреймворк echo,
(4 роута, разные методы http, разные способы передачи данных - by query, by body, etc.)
завернуть в докер контейнер, и запустить локально
проверить роуты апи через постман либо через запросы с другого го инстанса (net/http)
загрузить лабку на гитхаб/гитлаб
*/

func main() {
	db, err := sqlx.Open("postgres", "postgres://root:secret@localhost/books_api?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	store := Store.New(db)
	app := App.New(store)
	e := echo.New()

	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.GET("/books", app.GetBooks)
	e.GET("/books/:id", app.GetBookByID)
	e.POST("/books", app.SaveBook)
	e.PUT("/books/:id", app.UpdateBook)
	e.DELETE("/books/:id", app.DeleteBook)

	e.Logger.Fatal(e.Start(":8000"))
}
