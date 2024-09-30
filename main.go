package main

import (
	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-tech-blog/handler"
	"go-tech-blog/repository"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"os"
)

var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)

	// ホームパスにリクエストが来ると、 `articleIndex`関数が処理される
	e.GET("/", handler.ArticleIndex) // TOPページに記事一覧

	// 記事一覧画面には "/" と "/articles" の両方でアクセス可能
	e.GET("/articles", handler.ArticleIndex)           // 一覧画面
	e.GET("/articles/new", handler.ArticleNew)         // 新規作成画面
	e.GET("/articles/:articleID", handler.ArticleShow) // 詳細画面
	e.GET("/articles/:id/edit", handler.ArticleEdit)   // 編集画面

	// JSON返却処理は"/api"で開始する
	e.GET("/api/articles", handler.ArticleList)                 // 一覧
	e.POST("/api/articles", handler.ArticleCreate)              // 作成
	e.DELETE("/api/articles/:articleID", handler.ArticleDelete) // 削除

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
	e.Use(middleware.CSRF())    // CSRF対策

	// `src/css` ディレクトリ配下のファイルに `/css` のパスでアクセスできるようにする
	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	e.Validator = &CustomValidator{validator: validator.New()}

	// アプリインスタンスを返却
	return e
}

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
