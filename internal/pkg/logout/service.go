package logout

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
	Logout(c *context.Context, request *LogoutRequest) error
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

func (s *service) Logout(c *context.Context, request *LogoutRequest) error {
	timeNow := utils.TimeNowLocationTH()
	date := timeNow.Format("2006-01-02")
	time := timeNow.Format("15:04:05")
	userLog := &models.AlLogConnectModel{
		// Application: models.Application,
		// Log_name: c.GetEmployeeID(),
		// Log_type: models.Out,
		Log_type:    "logOut",
		Log_ip:      utils.GetIpAddress(),
		Log_level:   c.GetPermissions(),
		Log_user_id: c.GetEmployeeID(),
		Log_status:  "a",
		Log_date:    date,
		Log_time:    time,
	}
	// utils.PrintFormatJSON(userLog)
	err := s.libSQL.Create(c.GetTestDatabase(), userLog)
	if err != nil {
		logrus.Errorf("[Logout] Logout error: %s", err)
		return err
	}
	return nil
}
