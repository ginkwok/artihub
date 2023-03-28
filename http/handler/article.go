package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/ginkwok/artihub/model"
	"github.com/ginkwok/artihub/usecase"
	"github.com/ginkwok/artihub/util"
)

func GetAllArticlesHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)

	articles, err := usecase.GetAllArticles(ctx)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func CreateArticleHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)

	var article *model.Article
	err := c.BindJSON(&article)
	if err != nil {
		logger.Errorln(err)
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid request"},
		)
		return
	}
	if article.Title == "" {
		logger.Errorln("invalid article data")
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article data"},
		)
		return
	}

	article, err = usecase.CreateArticle(ctx, article)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

func GetArticleByIDHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)

	idstr := c.Param("article_id")
	if idstr == "" {
		logger.Errorln("invalid article id parameter")
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article id parameter"},
		)
		return
	}
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		logger.Errorln("invalid article id parameter", id, err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article id parameter," + err.Error()},
		)
		return
	}
	if id <= 0 {
		logger.Errorln("invalid article id parameter", id)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article id parameter"},
		)
		return
	}

	article, err := usecase.GetArticleByID(ctx, id)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

func UpdateArticleHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)

	idstr := c.Param("article_id")
	if idstr == "" {
		logger.Errorln("invalid article id parameter")
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article id parameter"},
		)
		return
	}
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		logger.Errorln("invalid article id parameter", id, err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article id parameter," + err.Error()},
		)
		return
	}
	if id <= 0 {
		logger.Errorln("invalid article id parameter", id)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article id parameter"},
		)
		return
	}

	var article *model.Article
	err = c.BindJSON(&article)
	if err != nil {
		logger.Errorln(err)
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid request"},
		)
		return
	}
	if article.Title == "" {
		logger.Errorln("invalid article data")
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article data"},
		)
		return
	}
	article.ID = uint(id)

	article, err = usecase.UpdateArticle(ctx, article)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, article)
}

func DeleteArticleByIDHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)

	idstr := c.Param("article_id")
	if idstr == "" {
		logger.Errorln("invalid article id parameter")
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article id parameter"},
		)
		return
	}
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		logger.Errorln("invalid article id parameter", id, err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article id parameter," + err.Error()},
		)
		return
	}
	if id <= 0 {
		logger.Errorln("invalid article id parameter", id)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid article id parameter"},
		)
		return
	}

	err = usecase.DeleteArticleByID(ctx, id)
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
