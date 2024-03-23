package route

import (
	"net/http"
	"test-echo/controllers"
	"test-echo/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! this is echo")
	})

	e.GET("/users", controllers.FecthAllUsers, middleware.IsAuth)
	e.POST("/users", controllers.CreateUsers)
	e.PUT("/users/:id", controllers.UpdateUsers)
	e.DELETE("/users/:id", controllers.DeleteUsers)
	e.GET("/users/:id", controllers.DetailUsers)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)

	e.POST("/login", controllers.CheckLogin)

	return e
}
