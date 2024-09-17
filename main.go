package main

import (
	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-tech-blog/handler"
	"go-tech-blog/repository"
	"log"
	"os"
)

var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)

	// ホームパスにリクエストが来ると、 `articleIndex`関数が処理される
	e.GET("/", handler.ArticleIndex)
	e.GET("/new", handler.ArticleNew)
	e.GET("/:id", handler.ArticleShow)
	e.GET("/:id/edit", handler.ArticleEdit)

	// Webサーバーをポート番号 8080 で起動する
	e.Logger.Fatal(e.Start(":8080"))
}

// MySQLに接続する
func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
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
