package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-test-simple/internal/handlers"
	"github.com/lailiseptiandi/go-test-simple/internal/middlewares"
	"github.com/lailiseptiandi/go-test-simple/internal/repository"
	"github.com/lailiseptiandi/go-test-simple/internal/services"
	"gorm.io/gorm"
)

func InitRoutes(dbs *gorm.DB, r *gin.Engine) {

	// heatlh check
	r.GET("health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "success check health")
	})

	// repository
	paymentRepo := repository.NewPaymentRepository(dbs)
	userRepo := repository.NewUserRepository(dbs)

	// service
	paymentService := services.NewPaymentService(&paymentRepo)
	userService := services.NewUserService(&userRepo)
	authService := services.NewAuthService(&userRepo)

	// handler
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	// api group
	apiGroupRoute := r.Group("api/v1/")

	// login register
	apiGroupRoute.POST("login", authHandler.Login)
	apiGroupRoute.POST("register", authHandler.Register)

	// User Route
	userGroupRoute := apiGroupRoute.Group("users")
	userGroupRoute.POST("/", middlewares.AuthRequired(), userHandler.Create)
	userGroupRoute.GET("/", middlewares.AuthRequired(), userHandler.Get)
	userGroupRoute.GET("/:id", middlewares.AuthRequired(), userHandler.GetByID)
	userGroupRoute.PUT("/:id", middlewares.AuthRequired(), userHandler.Update)
	userGroupRoute.DELETE("/:id", middlewares.AuthRequired(), userHandler.Delete)

	// Payment Route
	paymentGroupRoute := apiGroupRoute.Group("payment")
	paymentGroupRoute.POST("/", middlewares.AuthRequired(), paymentHandler.CreatePayment)
	paymentGroupRoute.POST("/idempotency_key", middlewares.AuthRequired(), paymentHandler.CreatePaymentIdempotencyKey)

}
