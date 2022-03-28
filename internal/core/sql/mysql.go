package sql

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	// MysqlPrpoDatabase Database global variable database
	StartDatabase = &gorm.DB{}
)

// Session session
type Session struct {
	Database *gorm.DB
}

// Configuration config mysql
type Configuration struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

// InitConnectionMysql open initialize a new db connection.
func InitConnectionMysql(config Configuration) (*Session, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DatabaseName,
	)

	database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return &Session{Database: database}, nil
}

// ConfigConnection config connection
type ConfigConnection struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

// NewConnectionMysql new db connection.
func NewConnectionMysql(cf ConfigConnection) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		cf.Username,
		cf.Password,
		cf.Host,
		cf.Port,
		cf.DatabaseName,
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		logrus.Errorf("[NewConnectionMysql] failed to connect to the database error: %s", err)
		return nil, err
	}

	return db, nil
}

// DebugMysqlPrpoDatabase set debug mysql prpo
// func MyDatabase() {
// 	TestDatabase = TestDatabase.Debug()
// }
func SDatabase(){
	StartDatabase = StartDatabase.Debug()
}


// StopConnectionMysql stop db connection.
func StopConnectionMysql(database *gorm.DB) error {
	//ใช้สำหรับไม่ให้ โพรเซส ค้าง
	sqlDB, err := database.DB()
	if err != nil {
		return err
	}
	_ = sqlDB.Close()

	return nil
}
