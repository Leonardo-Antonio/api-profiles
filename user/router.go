package user

import "github.com/labstack/echo"

func Router(s iuser, e *echo.Echo) {
	hand := newHandler(s)
	group := e.Group("/v1/users")
	group.GET("", hand.GetAll)
	group.POST("/signup", hand.SignUp)
	group.POST("/signin", hand.SignIn)
	group.DELETE("", hand.Delete)
	group.PUT("", hand.Update)
}
