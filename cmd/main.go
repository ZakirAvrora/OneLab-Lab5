package main

import (
	"ZakirAvrora/Lab4/src/App"
	"ZakirAvrora/Lab4/src/Store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*lab4:
создать template crud используя фреймворк echo,
(4 роута, разные методы http, разные способы передачи данных - by query, by body, etc.)
завернуть в докер контейнер, и запустить локально
проверить роуты апи через постман либо через запросы с другого го инстанса (net/http)
загрузить лабку на гитхаб/гитлаб
*/

func main() {

	store := Store.New("./storage/books.json")
	app := App.New(store)
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/books", app.GetBooks)
	e.GET("/books/:id", app.GetBookByID)
	e.POST("/books", app.SaveBook)
	e.PUT("/books/:id", app.UpdateBook)
	e.DELETE("/books/:id", app.DeleteBook)

	e.Logger.Fatal(e.Start(":8080"))
}
