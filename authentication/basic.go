package authentication

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

var (
	id = os.Getenv("ID")
	pw = os.Getenv("PW")
)

func BasicAuth() echo.MiddlewareFunc  {
	return middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == id && password == pw {
			return true,nil
		}
		return false,nil
	})
}