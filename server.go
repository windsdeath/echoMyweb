package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func show(c echo.Context) error {
	// Get team and member from the query string
	names := c.QueryParam("names")
	ages := c.QueryParam("ages")
	return c.String(http.StatusOK, "이름:"+names+", 나이:"+ages)
}
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/save", save)
	e.Logger.Fatal(e.Start(":8000"))

}
