package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/ginkwok/artihub/config"
	"github.com/ginkwok/artihub/database"
	"github.com/ginkwok/artihub/http/handler"
	"github.com/ginkwok/artihub/http/middleware"
	"github.com/ginkwok/artihub/util"
)

func main() {
	logger := util.NewLogger()
	defer logger.Sync()

	db := database.GetDB(
		viper.GetString("db.mysql.host"),
		viper.GetString("db.mysql.port"),
		viper.GetString("db.mysql.username"),
		viper.GetString("db.mysql.password"),
		viper.GetString("db.mysql.database"),
	)

	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(logger))
	router.Use(middleware.MySQLMiddleware(db))

	v1 := router.Group("/api/v1")
	{
		articles := v1.Group("/articles")
		{
			articles.GET("/", handler.GetAllArticlesHandler)
			articles.GET("/:id", handler.GetArticleByIDHandler)
			articles.POST("/", handler.CreateArticleHandler)
			articles.PUT("/:id", handler.UpdateArticleHandler)
			articles.DELETE("/:id", handler.DeleteArticleByIDHandler)
		}
		authors := v1.Group("/authors")
		{
			authors.GET("/", nil)
			authors.GET("/:id", nil)
			authors.POST("/", nil)
			authors.PUT("/:id", nil)
			authors.DELETE("/:id", nil)
		}
		tags := v1.Group("/tags")
		{
			tags.GET("/", nil)
			tags.GET("/:id", nil)
			tags.POST("/", nil)
			tags.PUT("/:id", nil)
			tags.DELETE("/:id", nil)
		}
	}
	router.Run(":" + viper.GetString("server.port"))
}
