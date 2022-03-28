package token

import (
	"time"

	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/GolfPhumrapee/finance/internal/core/utils"
	"github.com/GolfPhumrapee/finance/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

// Service employee service interface
type Service interface {
	Create(c *context.Context, uid, level string) (*models.Token, error)
}

type service struct {
	config *config.Configs
	result *config.ReturnResult
	// libSQL   lib.SQL
}

// NewService new employee service
func NewService() Service {
	return &service{
		config: config.CF,
		result: config.RR,
		// libSQL:   lib.NewSQL(),
	}
}

// Create create token
func (s *service) Create(c *context.Context, uid, level string) (*models.Token, error) {
	object, err := s.createToken(c, uid, level)
	if err != nil {
		logrus.Errorf("create jwt token error: %s", err)
		return nil, err
	}

	return object, nil
}

func (s *service) createToken(c *context.Context, uid, level string) (*models.Token, error) {
	claims := &context.Claims{}
	now := time.Now()
	claims.Subject = uid
	claims.Issuer = "prpo.ace-energy.co.th"
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = now.Add(s.getAccessExpireAt()).Unix()
	payload := make(map[string]interface{})
	payload["level"] = level
	claims.Permissions = payload["level"]

	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	accessToken, err := t.SignedString(utils.SignKey)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	claims.ExpiresAt = now.Add(s.getRefreshExpireAt()).Unix()
	r := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	refreshToken, err := r.SignedString(utils.SignKey)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// database := c.GetPrpoDatabase()
	// token := &models.AccessToken{}
	// if err := database.
	// 	Where("token_type = ? AND user_id = ? AND status = ?", models.Active, uid, models.Active).
	// 	Find(token).Error; err != nil {
	// 	return nil, err
	// }

	// expireTime := now.Add(s.getAccessExpireAt()).Format("2006-01-02 15:04:05")
	// if token.TokenAccess == "" {

	// 	timeNow := utils.TimeNowLocationTH()
	// 	dateTime := timeNow.Format("2006-01-02 15:04:05")

	// 	tokenModel := &models.AccessToken{
	// 		CreatedAt:    dateTime,
	// 		UserId:       uid,
	// 		TokenAccess:  accessToken,
	// 		TokenRefresh: refreshToken,
	// 		ExpireTime:   expireTime,
	// 		TokenType:    "1", //กรณี เผื่อมี firebase token ใช้เป็น type 2 หรืออื่นๆ
	// 		Status:       models.Active,
	// 	}
	// 	err = s.libSQL.Create(c.GetPrpoDatabase(), tokenModel)
	// 	if err != nil {
	// 		logrus.Errorf("[Login] create access token error: %s", err)
	// 		return nil, err
	// 	}

	// } else {
	// 	token.TokenAccess = accessToken
	// 	token.TokenRefresh = refreshToken
	// 	token.ExpireTime = expireTime
	// 	token.Status = models.Active
	// 	err = s.libSQL.Update(c.GetPrpoDatabase(), token)
	// 	if err != nil {
	// 		logrus.Errorf("[Login] update access token error: %s", err)
	// 		return nil, err
	// 	}
	// }

	return &models.Token{
		EmployeeID:   uid,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

/*
func (s *service) createRefreshToken(uid string) (string, error) {
	rts := fmt.Sprintf("%s%s", uid, time.Now().String())
	h := sha1.New()
	_, err := h.Write([]byte(rts))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}*/
