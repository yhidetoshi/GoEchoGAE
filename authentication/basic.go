package authentication

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	id = os.Getenv("ID")
	pw = os.Getenv("PW")
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == id && password == pw {
			return true, nil
		}
		return false, nil
	})
}
