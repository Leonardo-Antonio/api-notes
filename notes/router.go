package notes

import "github.com/labstack/echo"

func Router(s iNotes, c *echo.Echo) {
	handler := newHandNote(s)
	group := c.Group("/v1/notes")
	group.GET("/:id", handler.GetAllNote)
	group.POST("", handler.NewNote)
	group.PUT("", handler.UpdateNote)
	group.DELETE("", handler.DeleteNote)
}
