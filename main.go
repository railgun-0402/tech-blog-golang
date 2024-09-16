package main

import (
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const tmplPath = "src/template/"

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

	// `src/css` ディレクトリ配下のファイルに `/css` のパスでアクセスできるようにする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	// アプリインスタンスを返却
	return e
}

func articleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Hello,World!",
		"Now":     time.Now(),
	}
	return render(c, "article/index.html", data)
}

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	// pongo2はHTMLレンダリングの関数
	// Mustはエラーがあればパニックを発生させる
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	// htmlBlobをよび、生成されたHTMLをバイトデータとして受け取る
	b, err := htmlBlob(file, data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	// 成功時はHTMLデータを返却
	return c.HTMLBlob(http.StatusOK, b)
}
