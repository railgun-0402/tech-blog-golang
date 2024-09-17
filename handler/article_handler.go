package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"go-tech-blog/repository"

	"github.com/labstack/echo/v4"
)

// ArticleIndex ...
func ArticleIndex(c echo.Context) error {
	// 記事データの一覧を取得する
	articles, err := repository.ArticleList()
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"Message":  "Article Index",
		"Now":      time.Now(),
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
