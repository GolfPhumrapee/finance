package deletestatus

import (
	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/GolfPhumrapee/finance/internal/core/lib"
	"github.com/GolfPhumrapee/finance/internal/core/utils"
	"github.com/GolfPhumrapee/finance/internal/models"
	"github.com/sirupsen/logrus"
)

// Service service interface
type Service interface {
	DeleteStatus(c *context.Context, request *DeleteStatus) error
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

func (s *service) DeleteStatus(c *context.Context, request *DeleteStatus) error {
	timeNow := utils.TimeNowLocationTH()
	date := timeNow.Format("2006-01-02")
	time := timeNow.Format("15:04:05")
	if request.Config_id != "" && request.Group_id != "" {
		ConfComConstModel := models.ConfComConstModel{}
		err := c.GetTestDatabase().Model(&ConfComConstModel).Where("config_id=? && group_id=?", request.Config_id, request.Group_id).Updates(map[string]interface{}{"node_status": 0, "log_date": date, "log_time": time, "log_user": c.GetEmployeeID()}).Error
		if err != nil {
			logrus.Errorf("[Delete] delete error: %s", err)
			return err
		}
	}
	return nil

}
