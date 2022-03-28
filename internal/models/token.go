package models

import (
	"time"
)

// Token token model
type Token struct {
	AccessToken  string `json:"access_token" gorm:"-"`
	RefreshToken string `json:"refresh_token"`
	EmployeeID   string `json:"-"`
}

// RefreshToken refresh token model
type RefreshToken struct {
	RefreshToken string    `json:"refreshToken"`
	Used         bool      `json:"used"`
	ExpireTime   time.Time `json:"expireTime"`
	UserId       string    `json:"userId"`
}

type AccessToken struct {
	Id           string `json:"id" gorm:"primary_key"`
	CreatedAt    string `json:"createdAt"`
	UserId       string `json:"userId"`
	TokenAccess  string `json:"tokenAccess"`
	TokenRefresh string `json:"tokenRefresh"`
	TokenType    string `json:"tokenType"`
	ExpireTime   string `json:"expireTime"`
	Status       int32  `json:"status" `
}

// TableName overrides the table name used by AccessToken to `access_token`
func (AccessToken) TableName() string {
	return "access_token"
}
