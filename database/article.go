package database

import (
	"gorm.io/gorm"

	"github.com/ginkwok/artihub/model"
)

func GetAllArticles(db *gorm.DB) ([]*model.Article, error) {
	var articles []*model.Article
	result := db.Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func GetArticleByID(db *gorm.DB, id int64) (*model.Article, error) {
	var article *model.Article
	err := db.First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func CreateArticle(db *gorm.DB, article *model.Article) (*model.Article, error) {
	err := db.Select("Title", "Content").Create(&article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func UpdateArticle(db *gorm.DB, article *model.Article) error {
	oldarticle, err := GetArticleByID(db, int64(article.ID))
	if err != nil {
		return err
	}
	return db.Model(&oldarticle).Updates(model.Article{
		Title:   article.Title,
		Content: article.Content,
	}).Error
}

func DeleteArticleByID(db *gorm.DB, id int64) error {
	return db.Delete(&model.Article{}, id).Error
}
