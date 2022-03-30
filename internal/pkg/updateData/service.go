package updateData

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
	Update(c *context.Context, request *UpdateData) error
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

func (s *service) Update(c *context.Context, request *UpdateData) error {
	timeNow := utils.TimeNowLocationTH()
	date := timeNow.Format("2006-01-02")
	time := timeNow.Format("15:04:05")
	if request.Config_id != "" && request.Group_id != "" {
		ConfComConstModel := models.ConfComConstModel{}
		err := c.GetTestDatabase().Model(&ConfComConstModel).Where("config_id=? && group_id=?", request.Config_id, request.Group_id).Updates(map[string]interface{}{"create_date": date, "create_time": time, "create_user": request.Create_user, "node_name_th": request.Node_name_th, "node_name_en": request.Node_name_en, "node_desc": request.Node_desc, "node_url": request.Node_url, "node_ref_value": request.Node_ref_value, "node_ref_value2": request.Node_ref_value2, "node_sort": request.Node_sort, "node_status": request.Node_status, "log_status": request.Log_status, "log_date": date, "log_time": time, "log_user": c.GetEmployeeID()}).Error
		if err != nil {
			logrus.Errorf("[Update] Update error: %s", err)
			return err
		}
	}
	return nil

}

// update all
// 	timeNow := utils.TimeNowLocationTH()
// 	date := timeNow.Format("2006-01-02")
// 	time := timeNow.Format("15:04:05")
// 	userLog := &models.ConfComConstModel{
// 		Config_id:       request.Config_id,
// 		Group_id:        request.Group_id,
// 		Create_date:     request.Create_date,
// 		Create_time:     request.Create_time,
// 		Create_user:     request.Create_user,
// 		Node_name_th:    request.Node_name_th,
// 		Node_name_en:    request.Node_name_en,
// 		Node_desc:       request.Node_desc,
// 		Node_url:        request.Node_url,
// 		Node_ref_value:  request.Node_ref_value,
// 		Node_ref_value2: request.Node_ref_value2,
// 		Node_sort:       request.Node_sort,
// 		Node_status:     request.Node_status,
// 		Log_status:      request.Log_status,
// 		Log_date:        date,
// 		Log_time:        time,
// 		Log_user:        request.Log_user,
// 	}
// 	// utils.PrintFormatJSON(userLog)
// 	if request.Config_id != "" && request.Group_id != "" {
// 		err := s.libSQL.Update(c.GetTestDatabase().Debug().Where("config_id=? && group_id=?", request.Config_id, request.Group_id), userLog)
// 		if err != nil {
// 			logrus.Errorf("[update] create log update error: %s", err)
// 			return err
// 		}
// 	}
// 	return nil

// }
