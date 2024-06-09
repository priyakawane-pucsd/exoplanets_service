package exoplanets

import (
	"context"
	"exoplanetservice/logger"
	"exoplanetservice/models/dto"
	"exoplanetservice/models/filters"
	"exoplanetservice/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Config struct {
}

type Service interface {
	CreateExoplanets(ctx context.Context, req *dto.ExoplanetRequest) (*dto.Exoplanet, error)
	GetExoplanets(ctx context.Context, filter *filters.ExoplanetFilter, limit, offset int) (*dto.ListExoplanetResponse, error)
	GetExoplanetById(ctx context.Context, exoplanetId string) (*dto.ExoplanetByIdResponse, error)
	UpdateExoplanetById(ctx context.Context, id string, exoplanet dto.ExoplanetRequest) error
	DeleteExoplanetById(ctx context.Context, exoplanetId string) error
	CalculateFuelEstimation(ctx context.Context, exoplanetId string, crewCapacity int) (*dto.FuelEstimationResponse, error)
}

type Handler struct {
	Config  *Config
	service Service
}

func NewHandler(ctx context.Context, cfg *Config, service Service) *Handler {
	return &Handler{service: service, Config: cfg}
}

func (h *Handler) Register(router gin.IRouter) {
	router.POST("/exoplanets", h.CreateExoplanets)
	router.GET("/exoplanets", h.GetExoplanets)
	router.GET("/exoplanets/:id", h.GetExoplanetById)
	router.PUT("/exoplanets/:id", h.UpdateExoplanetById)
	router.DELETE("/exoplanets/:id", h.DeleteExoplanetById)
	router.GET("/exoplanets/:id/fuel-estimation", h.CalculateFuelEstimation)
}

// CreateExoplanets godoc
// @Summary Create a new exoplanet
// @Description Create a new exoplanet with the provided details
// @Tags Exoplanets
// @Accept json
// @Produce json
// @Param exoplanet body dto.ExoplanetRequest true "Exoplanet Request"
// @Success 201 {object} dto.Exoplanet "Created successfully"
// @Failure 400 {object} utils.CustomError "Invalid request body"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /exoplanetservice/exoplanets [post]
func (h *Handler) CreateExoplanets(ctx *gin.Context) {
	var req dto.ExoplanetRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		logger.Error(ctx, "Failed to parse request body: %s", err.Error())
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid request body"))
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		logger.Error(ctx, "Validation error: %s", err.Error())
		utils.WriteError(ctx, err)
		return
	}

	res, err := h.service.CreateExoplanets(ctx, &req)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}

// GetExoplanets godoc
// @Summary Get a list of exoplanets
// @Description Retrieve a paginated list of exoplanets
// @Tags Exoplanets
// @Accept json
// @Produce json
// @Param limit query int false "Limit the number of exoplanets returned" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Param radius query float64 false "Radius of the exoplanet" default(0.0)
// @Param mass query float64 false "Mass of the exoplanet" default(0.0)
// @Success 200 {object} dto.ListExoplanetResponse "List of exoplanets"
// @Failure 400 {object} utils.CustomError "Invalid request parameters"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /exoplanetservice/exoplanets [get]
func (h *Handler) GetExoplanets(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")
	radiusStr := ctx.DefaultQuery("radius", "0.0")
	massStr := ctx.DefaultQuery("mass", "0.0")

	// Parse limit into an integer
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid limit parameter"))
		return
	}

	// Parse offset into an integer
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid offset parameter"))
		return
	}
	var filter filters.ExoplanetFilter

	filter.Radius, err = strconv.ParseFloat(radiusStr, 64)
	if err != nil {
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid radius parameter"))
		return
	}

	filter.Mass, err = strconv.ParseFloat(massStr, 64)
	if err != nil {
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid mass parameter"))
		return
	}

	exoplanets, err := h.service.GetExoplanets(ctx, &filter, limit, offset)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, exoplanets)
}

// @Summary Get an exoplanet by ID
// @Description Retrieve detailed information about an exoplanet by its ID
// @Tags Exoplanets
// @Accept json
// @Produce json
// @Param id path string true "Exoplanet ID"
// @Success 200 {object} dto.ExoplanetByIdResponse "Successful response containing exoplanet details"
// @Failure 400 {object} utils.CustomError "Invalid request format"
// @Failure 404 {object} utils.CustomError "Exoplanet not found"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /exoplanetservice/exoplanets/{id} [get]
func (h *Handler) GetExoplanetById(ctx *gin.Context) {
	// Extract the exoplnet ID from the request context
	exoplanetId := ctx.Param("id")
	exoplanet, err := h.service.GetExoplanetById(ctx, exoplanetId)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, exoplanet)
}

// @Summary Update an exoplanet by ID
// @Description Update an exoplanet by its ID
// @Tags Exoplanets
// @Accept json
// @Produce json
// @Param id path string true "Exoplanet ID"
// @Param updateExoplanetRequest body dto.ExoplanetRequest true "Update Exoplanet Request"
// @Success 200 {string} string "Exoplanet updated successfully"
// @Failure 400 {object} utils.CustomError "Invalid request format"
// @Failure 404 {object} utils.CustomError "Exoplanet not found"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /exoplanetservice/exoplanets/{id} [put]
func (h *Handler) UpdateExoplanetById(ctx *gin.Context) {
	exoplanetId := ctx.Param("id")
	var updateRequest dto.ExoplanetRequest
	err := ctx.BindJSON(&updateRequest)
	if err != nil {
		logger.Error(ctx, "Failed to parse request body: %s", err.Error())
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid request body"))
		return
	}
	err = h.service.UpdateExoplanetById(ctx, exoplanetId, updateRequest)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, "Exoplanet updated successfully")
}

// @Summary Delete an exoplanet by ID
// @Description Delete an exoplanet by its ID
// @Tags Exoplanets
// @Accept json
// @Produce json
// @Param id path string true "Exoplanet ID"
// @Success 200 {string} string "Exoplanet deleted successfully"
// @Failure 400 {object} utils.CustomError "Invalid request format"
// @Failure 404 {object} utils.CustomError "Exoplanet not found"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /exoplanetservice/exoplanets/{id} [delete]
func (h *Handler) DeleteExoplanetById(ctx *gin.Context) {
	// Extract the exoplanet ID from the request context
	exoplanetId := ctx.Param("id")
	err := h.service.DeleteExoplanetById(ctx, exoplanetId)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, "Exoplanet deleted successfully")
}

// @Summary Calculate fuel estimation
// @Description Calculate fuel estimation for a trip to an exoplanet by its ID
// @Tags Exoplanets
// @Accept json
// @Produce json
// @Param id path string true "Exoplanet ID"
// @Param crewCapacity query int false "Crew capacity for the trip" default(0)
// @Success 200 {object} dto.FuelEstimationResponse "Successful response containing fuel estimation details"
// @Failure 400 {object} utils.CustomError "Invalid request format"
// @Failure 404 {object} utils.CustomError "Exoplanet not found"
// @Failure 500 {object} utils.CustomError "Internal server error"
// @Router /exoplanetservice/exoplanets/{id}/fuel-estimation [get]
func (h *Handler) CalculateFuelEstimation(ctx *gin.Context) {
	// Extract the id params from the request context
	exoplanetId := ctx.Param("id")
	// Extract the crewCapacity query from the request context
	crewCapacityStr := ctx.DefaultQuery("crewCapacity", "0")

	// Parse limit into an integer
	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		utils.WriteError(ctx, utils.NewBadRequestError("Invalid crewCapacity parameter"))
		return
	}

	res, err := h.service.CalculateFuelEstimation(ctx, exoplanetId, crewCapacity)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}
	utils.WriteResponse(ctx, res)
}
