package utils

import (
	"github.com/GolfPhumrapee/finance/internal/core/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// DatabaseConnect database connect
func DatabaseConnect(config sql.ConfigConnection) (*gorm.DB, error) {
	database, err := sql.NewConnectionMysql(config)
	if err != nil {
		logrus.Errorf("[newDatabase] new database connection error: %s", err)
		return nil, err
	}

	return database, nil
}
