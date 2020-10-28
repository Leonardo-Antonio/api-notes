package type_notes

import "github.com/labstack/echo"

func Router(s iTypeNote, c *echo.Echo)  {
	handTypeNote := newHandlerTN(s)

	group := c.Group("/v1/type-notes")
	group.GET("", handTypeNote.GetTypes)
	group.POST("", handTypeNote.NewActivity)
	group.DELETE("/:id", handTypeNote.DeleteType)
}
