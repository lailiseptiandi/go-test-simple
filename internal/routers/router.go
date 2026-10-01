package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-test-simple/internal/handler"
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

	// service
	paymentService := services.NewPaymentService(&paymentRepo)

	// handler
	paymentHandler := handler.NewPaymentHandler(paymentService)

	paymentGroupRoute := r.Group("payment")
	paymentGroupRoute.POST("/", paymentHandler.CreatePayment)
	paymentGroupRoute.POST("/idempotency_key", paymentHandler.CreatePaymentIdempotencyKey)
}
