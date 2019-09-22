package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/appengine"
	"net/http"

	// GAE
	"github.com/yhidetoshi/GoEchoGAE/handler"
	// local only
	//"./handler"
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
