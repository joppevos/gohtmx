package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("*.html")),
	}
}

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Initialize template rendering
	e.Renderer = newTemplate()
	users := []User{
		{Id: 1, Name: "Alice", Age: 30},
		{Id: 2, Name: "Bob", Age: 35},
		{Id: 3, Name: "Charlie", Age: 25},
	}

	// Serve the index.html page
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", users)
	})
	// delete user
	e.DELETE("/delete/:id", func(c echo.Context) error {
		id := c.Param("id")
		fmt.Println(id)
		// delete id from users
		for i, user := range users {
			if fmt.Sprint(user.Id) == id {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}
		fmt.Println(users)
		return c.Render(http.StatusOK, "users", nil)
	})
	// add user
	e.POST("/add", func(c echo.Context) error {
		name := c.FormValue("name")
		age := c.FormValue("age")
		fmt.Println(name, age)
		// add new user
		users = append(users, User{Id: len(users) + 1, Name: name, Age: 30})
		fmt.Println(users)
		return c.Render(http.StatusOK, "users", nil)
	})
	// Start the server
	e.Start(":8080")
}
