package ping

import (
	"context"
	"exoplanetservice/utils"

	"github.com/gin-gonic/gin"
)

type Config struct {
}

type Service interface {
	Ping(ctx context.Context) error
}
type Handler struct {
	cfg     *Config
	service Service
}

func NewHandler(ctx context.Context, cfg *Config, service Service) *Handler {
	return &Handler{
		cfg:     cfg,
		service: service,
	}
}

func (h *Handler) Register(router gin.IRouter) {
	router.GET("/ping", h.Ping)
}

// Ping godoc
// @Summary Ping the server
// @Description Check if the server is alive
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} string "Okay, I am alive!"
// @Failure 500 {object} error "Internal Server Error"
// @Router /exoplanetservice/ping [get]
func (h *Handler) Ping(ctx *gin.Context) {
	err := h.service.Ping(ctx)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteSuccess(ctx, "Okay, I am alive!")
}
