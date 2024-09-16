package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	// ホームパスにリクエストが来ると、 `articleIndex`関数が処理される
	e.GET("/", articleIndex)

	// Webサーバーをポート番号 8080 で起動する
	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	// Echoインスタンスの生成
	e := echo.New()

	// アプリに各種ミドルウェアを設定
	e.Use(middleware.Recover()) // パニック時、アプリが停止しないよう回復する
	e.Use(middleware.Logger())  // ログを記録する
	e.Use(middleware.Gzip())    // レスポンスをGzip圧縮して転送する

	// アプリインスタンスを返却
	return e
}

func articleIndex(c echo.Context) error {
	// StatusCode:200 文字列をレスポンス！
	return c.String(http.StatusOK, "Hello,World!")
}
