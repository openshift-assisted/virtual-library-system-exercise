package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"virtual-library/repositories"
	"virtual-library/services"

	"virtual-library/controllers"
	"virtual-library/docs"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookController := controllers.NewBookController(bookService)


	docs.SwaggerInfo.Title = "Book API"
	docs.SwaggerInfo.Description = "API for managing books"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.POST("", bookController.CreateBook)
			books.GET("", bookController.ListBooks)
			books.GET("/:id", bookController.GetBook)
			books.PUT("/:id", bookController.UpdateBook)
			books.DELETE("/:id", bookController.DeleteBook)
		}
	}

	return r
}
