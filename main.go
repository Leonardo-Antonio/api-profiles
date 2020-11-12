package main

import (
	"github.com/Leonardo-Antonio/api-profiles/dbutil"
	"github.com/Leonardo-Antonio/api-profiles/user"
	"github.com/labstack/echo"
)

func main() {
	pool := dbutil.GetConnection(dbutil.MYSQL)
	userStorage := user.NewStorage(pool)

	e := echo.New()

	user.Router(userStorage, e)
	e.Logger.Fatal(e.Start(":8080"))
}
