package Routes

import (
	"belajargo/Controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

// setup for routing
func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Belajar Go Rest Api")
	})
	e.GET("/course", Controllers.GetCourse)
	e.POST("/course", Controllers.CreateCourse)
	e.PUT("/course/:id", Controllers.UpdateCourse)
	e.DELETE("/course/:id", Controllers.DeleteCourse)
}
