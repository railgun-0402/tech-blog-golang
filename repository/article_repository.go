package repository

import (
	"database/sql"
	"go-tech-blog/model"
	"math"
	"time"
)

// ArticleList
// articleテーブルから一覧データを取得する/*
//func ArticleList() ([]*model.Article, error) {
//	query := `SELECT * FROM articles;`
//
//	// DBから取得した値を格納する変数を宣言
//	var articles []*model.Article
//
//	// Query を実行して、取得した値を変数に格納
//	if err := db.Select(&articles, query); err != nil {
//		return nil, err
//	}
//
//	return articles, nil
//}

// ArticleListByCursor ...
func ArticleListByCursor(cursor int) ([]*model.Article, error) {
	// 引数で渡されたカーソルの値が0以下の場合は、代わりにint型の最大値で置き換える
	if cursor <= 0 {
		cursor = math.MaxInt32
	}

	// 降順 & 10件取得
	query := `SELECT *
	FROM articles
	WHERE id < ?
	ORDER BY id DESC
	LIMIT 10`

	// スライス初期化
	articles := make([]*model.Article, 0, 10)

	if err := db.Select(&articles, query, cursor); err != nil {
		return nil, err
	}
	return articles, nil
}

// ArticleCreate 記事をDBに登録する
func ArticleCreate(article *model.Article) (sql.Result, error) {
	now := time.Now()

	article.Created = now
	article.Updated = now

	query := `INSERT INTO articles (title, body, created, updated) VALUES (:title, :body, :created, :updated);`

	// トランザクション開始
	tx := db.MustBegin()

	res, err := tx.NamedExec(query, article)
	if err != nil {
		tx.Rollback()

		return nil, err
	}
	tx.Commit()
	return res, nil
}
