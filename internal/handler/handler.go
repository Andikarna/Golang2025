package handler

import (
	"encoding/json"
	"net/http"

	"fundamental/internal/model"
    "fundamental/internal/repository"
)

var _ = model.Article{}

// GetArticlesHandler godoc
// @Summary Get all articles
// @Description Get details of all available articles
// @Tags Articles
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Article
// @Router /articles [get]
func GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
    articles, err := repository.GetAllArticles()
    if err != nil {
        http.Error(w, "Error fetching articles", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(articles)
}

// CreateArticleHandler godoc
// @Summary Create a new article
// @Description Create a new article with title, description and content
// @Tags Articles
// @Accept  json
// @Produce  json
// @Param article body model.Article true "New Article"
// @Success 200 {object} model.Article
// @Router /articles [post]
func CreateArticleHandler(w http.ResponseWriter, r *http.Request) {
    var article model.Article
    _ = json.NewDecoder(r.Body).Decode(&article)

    if err := repository.CreateArticle(&article); err != nil {
        http.Error(w, "Error saving article", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(article)
}
