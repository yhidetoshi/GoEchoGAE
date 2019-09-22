package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/appengine"
	"net/http"

	// GAE
	"github.com/yhidetoshi/apiEchoGAE/handler"
	"github.com/yhidetoshi/apiEchoGAE/authentication"
	// local only
	//"./handler"
	//"./authentication"
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
	e.GET("/metal", handler.FetchMetal())

	// サーバー起動 at local
	//e.Start(":1323")

	// GAE
	appengine.Main()
}

// GAE
func createMux() *echo.Echo {
	e := echo.New()
	return e
}