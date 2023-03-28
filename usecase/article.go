package usecase

import (
	"context"

	"go.uber.org/zap"

	"github.com/ginkwok/artihub/model"
	"github.com/ginkwok/artihub/service"
	"github.com/ginkwok/artihub/util"
)

func GetAllArticles(ctx context.Context) ([]*model.Article, error) {
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)
	articles, err := service.GetAllArticles(ctx)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return articles, nil
}
