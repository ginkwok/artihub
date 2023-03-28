package database

import (
	"gorm.io/gorm"

	"github.com/ginkwok/artihub/model"
)

func GetAllArticles(db *gorm.DB) ([]*model.Article, error) {
	var users []*model.Article
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
