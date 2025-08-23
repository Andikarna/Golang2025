package repository

import (
	"fundamental/internal/database"
	"fundamental/internal/model"
)

func GetAllArticles() ([]model.Article, error) {
    var articles []model.Article
    result := database.DB.Find(&articles)
    return articles, result.Error
}

func CreateArticle(article *model.Article) error {
    return database.DB.Create(article).Error
}