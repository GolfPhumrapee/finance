package insertData

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
	InsertData(c *context.Context, request *AddInformationRequest) error
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

func (s *service) InsertData(c *context.Context, request *AddInformationRequest) error {
	timeNow := utils.TimeNowLocationTH()
	date := timeNow.Format("2006-01-02")
	time := timeNow.Format("15:04:05")
	userLog := &models.ConfComConstModel{
		Config_id:       request.Config_id,
		Group_id:        request.Group_id,
		Create_date:     date,
		Create_time:     time,
		Create_user:     request.Create_user,
		Node_name_th:    request.Node_name_th,
		Node_name_en:    request.Node_name_en,
		Node_desc:       request.Node_desc,
		Node_url:        request.Node_url,
		Node_ref_value:  request.Node_ref_value,
		Node_ref_value2: request.Node_ref_value2,
		Node_sort:       request.Node_sort,
		Node_status:     request.Node_status,
		Log_status:      request.Log_status,
		Log_date:        date,
		Log_time:        time,
		Log_user:        request.Log_user,
	}
	// utils.PrintFormatJSON(userLog)
	if request.Create_user != "" && request.Node_status == 1 {
		err := s.libSQL.Create(c.GetTestDatabase(), userLog)
		if err != nil {
			logrus.Errorf("[insert] insert error: %s", err)
			return err
		}
	}
	return nil

}
