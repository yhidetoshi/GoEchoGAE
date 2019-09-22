package main

import (
	"./handler"
	"github.com/labstack/echo"
)

func main() {

	// Echoインスタンス作成
	e := echo.New()


	// ルーティング
	e.GET("/metal", handler.FetchMetal())

	// サーバー起動
	e.Start(":1323")
}