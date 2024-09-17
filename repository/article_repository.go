package repository

import "go-tech-blog/model"

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
