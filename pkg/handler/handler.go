package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pla1no/web-server-test/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	productRouter := router.Group("/product")
	{
		productRouter.GET("/", h.getAllProducts)
		productRouter.GET("/:id", h.getProductByID)
		productRouter.POST("/", h.createProduct)
		productRouter.PUT("/:id", h.updateProduct)
		productRouter.DELETE("/:id", h.deleteProduct)
	}

	measureRouter := router.Group("/measure")
	{
		measureRouter.GET("/", h.getAllMeasures)
		measureRouter.GET("/:id", h.getMeasureByID)
		measureRouter.POST("/", h.createMeasure)
		measureRouter.PUT("/:id", h.updateMeasure)
		measureRouter.DELETE("/:id", h.deleteMeasure)
	}
	return router
}
