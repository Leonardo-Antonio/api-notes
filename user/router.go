package user

import "github.com/labstack/echo"

func Router(storage iuser ,c *echo.Echo)  {
	handUser := newHandler(storage)

	group := c.Group("/v1/users")
	group.POST("/new", handUser.SignUp)
	group.POST("", handUser.LogIn)
}
