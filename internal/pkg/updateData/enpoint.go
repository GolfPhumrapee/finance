package updateData

import (
	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	Update(c *fiber.Ctx) error
}

type endpoint struct {
	config  *config.Configs
	result  *config.ReturnResult
	service Service
}

// NewEndpoint new migrate endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		config:  config.CF,
		result:  config.RR,
		service: NewService(),
	}
}

// Update
// @Tags Data
// @Summary Update
// @Description testswager
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request body UpdateData false "request body"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Security ApiKeyAuth
// @Router /Data/UpDateData [post]
func (ep *endpoint) Update(c *fiber.Ctx) error {
	return handlers.ResponseSuccess(c, ep.service.Update, &UpdateData{})
}
