package repository

import (
	"database/sql"
	"go-tech-blog/model"
	"time"
)

// ArticleList
// articleテーブルから一覧データを取得する/*
func ArticleList() ([]*model.Article, error) {
	query := `SELECT * FROM articles;`

	// DBから取得した値を格納する変数を宣言
	var articles []*model.Article

	// Query を実行して、取得した値を変数に格納
	if err := db.Select(&articles, query); err != nil {
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
