package delete

import (
	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/GolfPhumrapee/finance/internal/core/lib"
	"github.com/GolfPhumrapee/finance/internal/models"
	"github.com/sirupsen/logrus"
)

// Service service interface
type Service interface {
	Delete(c *context.Context, request *Delete) error
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

func (s *service) Delete(c *context.Context, request *Delete) error {
	userLog := &models.ConfComConstModel{
		Config_id: request.Config_id,
		Group_id:  request.Group_id,
	}
	if request.Config_id != "" && request.Group_id != "" {
		err := s.libSQL.Delete(c.GetTestDatabase().Debug().Where("create_user=?", request.Config_id), userLog)
		if err != nil {
			logrus.Errorf("[delete] create log delete error: %s", err)
			return err
		}
	}
	return nil

}
