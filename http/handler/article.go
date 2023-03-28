package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/ginkwok/artihub/usecase"
	"github.com/ginkwok/artihub/util"
)

func GetAllArticlesHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := ctx.Value(util.LoggerKey).(*zap.SugaredLogger)

	articles, err := usecase.GetAllArticles(ctx)
	if err != nil {
		logger.Errorln("invalid name parameter")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid name parameter"})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func GetArticleByIDHandler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func CreateArticleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func UpdateArticleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func DeleteArticleByIDHandler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
