package selectGroup

import (
	"github.com/GolfPhumrapee/finance/internal/core/lib"
	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/GolfPhumrapee/finance/internal/models"
)

// Service service interface
type Service interface {
	SelectGroup(c *context.Context, request *SelectGroup) ([]*models.ConfComConstModel, error)
}

type service struct {
	config *config.Configs
	result *config.ReturnResult
	libSQL lib.SQL
}

// NewService new service
func NewService() Service {
	return &service{
		config: config.CF,
		result: config.RR,
		libSQL: lib.NewSQL(),
	}
}

func (s *service) SelectGroup(c *context.Context, request *SelectGroup) ([]*models.ConfComConstModel, error) {
	response := []*models.ConfComConstModel{}
	database := c.GetTestDatabase().Debug()
	aa := database.Where("group_id=?", request.Group_id)
	bb := database.Where("node_status=?", 1)
	if request.Group_id != "" && request.Node_name_th != "" && request.Node_name_en != "" && request.Node_desc != "" && request.Node_url != "" && request.Node_ref_value != "" && request.Node_ref_value2 != "" {
		aa.Where("group_id=?", request.Group_id)
	} else {
		bb.Where("node_status=?", 1)
	}
	if request.Group_id != "" && request.Node_name_th != "" && request.Node_name_en != "" && request.Node_desc != "" && request.Node_url != "" && request.Node_ref_value != "" && request.Node_ref_value2 != "" {
		err := aa.Find(&response).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := bb.Find(&response).Error
		if err != nil {
			return nil, err
		}
	}
	return response, nil

}
