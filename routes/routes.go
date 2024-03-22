package route

import (
	"net/http"
	"test-echo/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! this is echo")
	})

	e.GET("/users", controllers.FecthAllUsers)
	e.POST("/users", controllers.CreateUsers)

	return e
}
