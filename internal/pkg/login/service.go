package login

import (
	"gitee.com/beshy/php2go"
	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/GolfPhumrapee/finance/internal/core/lib"
	"github.com/GolfPhumrapee/finance/internal/core/utils"
	"github.com/GolfPhumrapee/finance/internal/models"
	"github.com/GolfPhumrapee/finance/internal/pkg/token"
	"github.com/sirupsen/logrus"
)

// Service service interface
type Service interface {
	// decrytEncapsulate(encoding string) (string, error)
	Login(c *context.Context, request *LoginRequest) (map[string]interface{}, error)
}

type service struct {
	config   *config.Configs
	result   *config.ReturnResult
	tokenSrv token.Service
	libSQL   lib.SQL
}

// NewService new service
func NewService() Service {
	return &service{
		config:   config.CF,
		result:   config.RR,
		tokenSrv: token.NewService(),
		libSQL:   lib.NewSQL(),
	}
}

//encodeUser encode user
func (s *service) decrytEncapsulate(encoding string) (string, error) {
	keyIndex := s.config.KeyMypass
	key := utils.MD5HashHex(keyIndex)

	decode, err := decoding(encoding, key)
	if err != nil {
		return "", err
	}
	return decode, nil
}

func decoding(encode, key string) (string, error) {
	j := 0
	hash := ""
	key = php2go.Sha1(key)
	strLen := len(encode)
	keyLen := len(key)
	for i := 0; i < strLen; i += 2 {
		substr := php2go.Substr(encode, uint(i), 2)
		strrev := php2go.Strrev(substr)
		convert, err := php2go.BaseConvert(strrev, 36, 16)
		if err != nil {
			return "", err
		}
		hex, err := php2go.Hexdec(convert)
		if err != nil {
			return "", err
		}
		if j == keyLen {
			j = 0
		}
		ordKey := php2go.Ord(php2go.Substr(key, uint(j), 1))
		j++
		hash += php2go.Chr(int(hex) - ordKey)
	}
	return hash, nil
}

//ส่วนเรียกใช้งาน
func (s *service) Login(c *context.Context, request *LoginRequest) (map[string]interface{}, error) {
	arrReturn := make(map[string]interface{}) //สร้าง array โดยกำหนด string เป็น index ของ array
	user, err := s.decrytEncapsulate(request.Encode)
	userLogin := php2go.Explode("##XXX##", user)
	UserId := userLogin[0]
	Level := userLogin[1]
	user_status := 1
	response := &models.AlComUserModel{}
	database := c.GetTestDatabase()
	DB := database.Debug()
	if request.Encode != "" {
		DB = DB.Where("user_id=? && level_priv=? && user_status=?", UserId, Level, user_status)
	}
	err_1 := DB.Find(&response).Error
	if err_1 != nil {
		return nil, err
	}
	if response.User_id != "" {
		token, err := s.tokenSrv.Create(c, response.Applicant_id, response.Level_priv)
		if err != nil {
			logrus.Errorf("[Login] generate access token error: %s", err)
			return nil, err
		}
		// println("------------------------------")
		// println("old token :", token.AccessToken)
		arrReturn["User"] = response
		arrReturn["Token"] = token.AccessToken // สร้างและเก็บโทเค้นไว้ใน arrReturn
	}
	if response.User_id != "" {
		timeNow := utils.TimeNowLocationTH()
		date := timeNow.Format("2006-01-02")
		time := timeNow.Format("15:04:05")
		userLog := &models.AlLogConnectModel{
			// Application: models.Application,
			// Log_name: c.GetEmployeeID(),
			// Log_type: models.Out,
			Log_type:  "login",
			Log_ip:    utils.GetIpAddress(),
			Log_level: c.GetPermissions(),
			// Log_level: Level,
			Log_user_id: c.GetEmployeeID(),
			// Log_user_id: UserId,
			Log_status: "a",
			Log_date:   date,
			Log_time:   time,
		}
		// utils.PrintFormatJSON(userLog)
		err := s.libSQL.Create(c.GetTestDatabase(), userLog)
		if err != nil {
			logrus.Errorf("[Login] login error: %s", err)
			return nil, err
		}
	}
	return arrReturn, nil
}
