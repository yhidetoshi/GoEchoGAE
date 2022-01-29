package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yhidetoshi/apiEchoGAE/authentication"
	"google.golang.org/appengine"

	// GAE
	"github.com/yhidetoshi/apiEchoGAE/handler"
)

// GAE
var e = createMux()

func main() {

	// Echoインスタンス作成
	e := echo.New()
	// GAE
	http.Handle("/", e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	// Basic Auth
	e.Use(authentication.BasicAuth())

	// ルーティング
	e.GET("/metal", handler.FetchMetal)

	// サーバー起動 at local
	e.Start(":1323")

	// GAE
	appengine.Main()
}

// GAE
func createMux() *echo.Echo {
	e := echo.New()
	return e
}
