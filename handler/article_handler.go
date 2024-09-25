package handler

import (
	"go-tech-blog/model"
	"net/http"
	"strconv"
	"time"

	"go-tech-blog/repository"

	"github.com/labstack/echo/v4"
)

type ArticleCreateOutput struct {
	Article          *model.Article
	Message          string
	ValidationErrors []string
}

// ArticleCreate ...
func ArticleCreate(c echo.Context) error {
	// 送信フォーム内容を格納
	var article model.Article

	// レスポンスとして返却する内容
	var out ArticleCreateOutput

	// フォームの内容を構造体に埋め込み
	if err := c.Bind(&article); err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusBadRequest, out)
	}

	// バリデーション
	if err := c.Validate(&article); err != nil {
		c.Logger().Error(err.Error())

		// エラー内容をレスポンス内容に入れる
		out.ValidationErrors = article.ValidationErrors(err)

		// 許可されてないパラメータの場合は422
		return c.JSON(http.StatusUnprocessableEntity, out)
	}

	// repositoryを呼び出して、保存処理実行
	res, err := repository.ArticleCreate(&article)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.JSON(http.StatusInternalServerError, out)
	}

	// SQL 実行結果から作成されたレコードの ID を取得
	id, _ := res.LastInsertId()

	// article構造体にIDをセットする
	article.ID = int(id)
	out.Article = &article

	return c.JSON(http.StatusOK, out)
}

// ArticleIndex ...
func ArticleIndex(c echo.Context) error {
	// 記事データの一覧を取得する
	articles, err := repository.ArticleListByCursor(0)
	if err != nil {
		c.Logger().Error(err.Error())

		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Articles": articles, // 記事データをテンプレートエンジンに渡す
	}
	return render(c, "article/index.html", data)
}

// ArticleNew ...
func ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return render(c, "article/new.html", data)
}

// ArticleShow ...
func ArticleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Show",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/show.html", data)
}

// ArticleEdit ...
func ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/edit.html", data)
}
