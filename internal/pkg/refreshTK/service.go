package refreshTK

import (
	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/GolfPhumrapee/finance/internal/models"
	"github.com/GolfPhumrapee/finance/internal/pkg/token"
	"github.com/sirupsen/logrus"
)

// Service service interface
type Service interface {
	refreshTK(c *context.Context, request *LoginRequest) (map[string]interface{}, error)
}

type service struct {
	config   *config.Configs
	result   *config.ReturnResult
	tokenSrv token.Service
}

// NewService new service
func NewService() Service {
	return &service{
		config:   config.CF,
		result:   config.RR,
		tokenSrv: token.NewService(),
	}
}

func (s *service) refreshTK(c *context.Context, request *LoginRequest) (map[string]interface{}, error) {
	arrReturn := make(map[string]interface{})
	response := []*models.AlComUserModel{}
	database := c.GetTestDatabase()
	DB := database.Debug()
	if request.User_id != "" {
		DB.Where("user_id=?", request.User_id)
	}
	err := DB.Find(&response).Error
	if err != nil {
		return nil, err
	}
	if c.GetEmployeeID() != "" {
		token, err := s.tokenSrv.Create(c, c.GetEmployeeID(), c.GetPermissions())
		if err != nil {
			logrus.Errorf("[refreshTK] refresh token error: %s", err)
			return nil, err
		}
		arrReturn["RefreshToken"] = token.RefreshToken
	}
	return arrReturn, nil
}
