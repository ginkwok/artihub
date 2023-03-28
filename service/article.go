package service

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ginkwok/artihub/database"
	"github.com/ginkwok/artihub/model"
	"github.com/ginkwok/artihub/util"
)

func GetAllArticles(ctx context.Context) ([]*model.Article, error) {
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)
	db := ctx.Value(util.MySQLKey).(*gorm.DB)

	articles, err := database.GetAllArticles(db)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return articles, nil
}

func GetArticleByID(ctx context.Context, id int64) (*model.Article, error) {
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)
	db := ctx.Value(util.MySQLKey).(*gorm.DB)

	article, err := database.GetArticleByID(db, id)
	if err != nil {
		logger.Errorln(id, err)
		return nil, err
	}
	return article, nil
}

func CreateArticle(ctx context.Context, article *model.Article) (*model.Article, error) {
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)
	db := ctx.Value(util.MySQLKey).(*gorm.DB)

	article, err := database.CreateArticle(db, article)
	if err != nil {
		logger.Errorln(article, err)
		return nil, err
	}
	return article, nil
}

func UpdateArticle(ctx context.Context, article *model.Article) error {
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)
	db := ctx.Value(util.MySQLKey).(*gorm.DB)

	err := database.UpdateArticle(db, article)
	if err != nil {
		logger.Errorln(article, err)
		return err
	}
	return nil
}

func DeleteArticleByID(ctx context.Context, id int64) error {
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)
	db := ctx.Value(util.MySQLKey).(*gorm.DB)

	err := database.DeleteArticleByID(db, id)
	if err != nil {
		logger.Errorln(id, err)
		return err
	}
	return nil
}
