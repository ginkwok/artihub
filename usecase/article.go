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

func GetArticleByID(ctx context.Context, id int64) (*model.Article, error) {
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)
	article, err := service.GetArticleByID(ctx, id)
	if err != nil {
		logger.Errorln(id, err)
		return nil, err
	}
	return article, nil
}

func CreateArticle(ctx context.Context, article *model.Article) (*model.Article, error) {
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)

	article1, err := service.CreateArticle(ctx, article)
	if err != nil {
		logger.Errorln(article, err)
		return nil, err
	}

	article2, err := service.GetArticleByID(ctx, int64(article1.ID))
	if err != nil {
		logger.Errorln(article1, err)
		return nil, err
	}
	return article2, nil
}

func UpdateArticle(ctx context.Context, article *model.Article) (*model.Article, error) {
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)

	err := service.UpdateArticle(ctx, article)
	if err != nil {
		logger.Errorln(article, err)
		return nil, err
	}

	article1, err := service.GetArticleByID(ctx, int64(article.ID))
	if err != nil {
		logger.Errorln(article, err)
		return nil, err
	}
	return article1, nil
}

func DeleteArticleByID(ctx context.Context, id int64) error {
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)
	err := service.DeleteArticleByID(ctx, id)
	if err != nil {
		logger.Errorln(id, err)
		return err
	}
	return nil
}
