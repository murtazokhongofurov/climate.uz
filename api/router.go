package api

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	swaggerfile "github.com/swaggo/files"
	_ "gitlab.com/climate.uz/api/docs"
	v1 "gitlab.com/climate.uz/api/handler/v1"
	"gitlab.com/climate.uz/api/middleware"
	"gitlab.com/climate.uz/api/tokens"
	"gitlab.com/climate.uz/config"
	"gitlab.com/climate.uz/internal/controller/storage"
	"gitlab.com/climate.uz/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Options struct {
	Conf           config.Config
	Logger         logger.Logger
	Storage        storage.StorageI
	CasbinEnforcer *casbin.Enforcer
}

// New ...
// @title           Climate_uz api
// @version         1.0
// @description     This is climate_uz server api server
// @termsOfService  2 term climate_uz

// @contact.name   	Murtazoxon
// @contact.url    	https://t.me/murtazokhon_gofurov
// @contact.email  	gofurovmurtazoxon@gmail.com

// @host      		localhost:7070
// @BasePath  		/v1

// @securityDefinitions.apikey BearerAuth
// @in 		header
// @name 	Authorization
func New(option Options) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	jwtHandler := tokens.JWTHandler{
		SigninKey: option.Conf.SigninKey,
		Log:       option.Logger,
	}
	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:     option.Logger,
		Storage:    option.Storage,
		Cfg:        option.Conf,
		JwtHandler: jwtHandler,
	})

	router.Use(middleware.NewAuth(option.CasbinEnforcer, jwtHandler, config.Load()))

	router.MaxMultipartMemory = 8 << 20 // 8 Mib

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "App is running",
		})
	})
	api := router.Group("/v1")

	// Category apis
	api.POST("/category", handlerV1.CreateCategory)
	api.GET("/category/:id", handlerV1.GetCategory)
	api.PATCH("/category", handlerV1.EditCategoryName)
	api.GET("/categories", handlerV1.GetAllGategory)
	api.DELETE("category/:id", handlerV1.DeleteCategory)

	// User apis
	api.POST("/user", handlerV1.CreateUser)
	api.GET("/user/:id", handlerV1.GetUser)
	api.GET("/users", handlerV1.GetAllUsers)
	api.PATCH("/user", handlerV1.UpdateUser)
	api.DELETE("/user/:id", handlerV1.DeleteUser)

	// Product apis
	api.POST("/product", handlerV1.CreateProduct)
	api.GET("/product/:id", handlerV1.GetProductInfo)
	api.PATCH("/product", handlerV1.UpdateProduct)
	api.GET("/products", handlerV1.GetAllProduct)
	api.DELETE("/product/:id", handlerV1.DeleteProductInfo)
	api.DELETE("/products/:id", handlerV1.DeleteProductsCategoryId)

	// News product apis
	api.POST("/news", handlerV1.CreateNewsProduct)
	api.GET("/news/:id", handlerV1.GetNewsProduct)
	api.GET("/news/incategory/:id", handlerV1.GetNewsByCategoryId)
	api.GET("/allnews", handlerV1.GetAllNewsProducts)
	api.PATCH("/news", handlerV1.UpdateNewsProduct)
	api.DELETE("/news/:id", handlerV1.DeleteNewsProduct)
	api.DELETE("news/category/:id", handlerV1.DeleteNewsByCategoryId)

	api.GET("/admin/:username/:password", handlerV1.AdminLogin)
	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfile.Handler, url))

	return router
}
