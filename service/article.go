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
