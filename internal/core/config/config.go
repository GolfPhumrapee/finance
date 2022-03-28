package config

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// CF -> for use configs model
	CF = &Configs{}
)

// DatabaseConfig database config model
type DatabaseConfig struct {
	Host         string `mapstructure:"HOST"`
	Port         int    `mapstructure:"PORT"`
	Username     string `mapstructure:"USERNAME"`
	Password     string `mapstructure:"PASSWORD"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
	DriverName   string `mapstructure:"DRIVER_NAME"`
	Timeout      string `mapstructure:"TIMEOUT"`
}

type ExpireTime struct {
	Day    time.Duration `mapstructure:"DAY"`
	Hour   time.Duration `mapstructure:"HOUR"`
	Minute time.Duration `mapstructure:"MINUTE"`
}

// Configs config models
type Configs struct {
	UniversalTranslator *ut.UniversalTranslator
	Validator           *validator.Validate
	KeyMypass           string `mapstructure:"key_mypass"`
	// configminio
	Minio struct {
		MINIO_ENDPOINT  string `mapstructure:"MINIO_ENDPOINT"`
		MINIO_PORT      string `mapstructure:"MINIO_PORT"`
		MINIO_ACCESSKEY string `mapstructure:"MINIO_ACCESSKEY"`
		MINIO_SECRETKEY string `mapstructure:"MINIO_SECRETKEY"`
		MINIO_BUCKET    string `mapstructure:"MINIO_BUCKET"`
		MINIO_PATH      string `mapstructure:"MINIO_PATH"`
		UseSSL          bool   `mapstructure:"UseSSL"`
	}
	App struct {
		ProjectID  string `mapstructure:"PROJECT_ID"`
		SystemName string `mapstructure:"SYSTEM_NAME"`
		WebBaseURL string `mapstructure:"WEB_BASE_URL"`
		APIBaseURL string `mapstructure:"API_BASE_URL"`
		Version    string `mapstructure:"VERSION"`
		Release    bool   `mapstructure:"RELEASE"`
		Port       int    `mapstructure:"PORT"`
	} `mapstructure:"APP"`
	HTTPServer struct {
		ReadTimeout  time.Duration `mapstructure:"READ_TIMEOUT"`
		WriteTimeout time.Duration `mapstructure:"WRITE_TIMEOUT"`
		IdleTimeout  time.Duration `mapstructure:"IDLE_TIMEOUT"`
	} `mapstructure:"HTTP_SERVER"`
	Swagger struct {
		Title       string   `mapstructure:"TITLE"`
		Version     string   `mapstructure:"VERSION"`
		Host        string   `mapstructure:"HOST"`
		Description string   `mapstructure:"DESCRIPTION"`
		Schemes     []string `mapstructure:"SCHEMES"`
		Enable      bool     `mapstructure:"ENABLE"`
	} `mapstructure:"SWAGGER"`
	Mysql struct {
		Sql struct {
			Host         string `mapstructure:"HOST"`
			Port         int    `mapstructure:"PORT"`
			Username     string `mapstructure:"USERNAME"`
			Password     string `mapstructure:"PASSWORD"`
			DatabaseName string `mapstructure:"DATABASE_NAME"`
			DriverName   string `mapstructure:"DRIVER_NAME"`
		} `mapstructure:"Sql"`
	} `mapstructure:"MYSQL"`
	JWT struct {
		Access struct {
			ExpireTime `mapstructure:"EXPIRE_TIME"`
		} `mapstructure:"ACCESS"`
		Refresh struct {
			ExpireTime `mapstructure:"EXPIRE_TIME"`
		} `mapstructure:"REFRESH"`
		PrivateKeyPath string `mapstructure:"PRIVATE_KEY_PATH"`
		PublicKeyPath  string `mapstructure:"PUBLIC_KEY_PATH"`
	} `mapstructure:"JWT"`
}

// InitConfig init config
func InitConfig(configPath string) error {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName("config")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		logrus.Error("read config file error:", err)
		return err
	}

	if err := bindingConfig(v, CF); err != nil {
		logrus.Error("binding config error:", err)
		return err
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := bindingConfig(v, CF); err != nil {
			logrus.Error("binding error:", err)
			return
		}
	})

	return nil
}

// bindingConfig binding config
func bindingConfig(vp *viper.Viper, cf *Configs) error {
	if err := vp.Unmarshal(&cf); err != nil {
		logrus.Error("unmarshal config error:", err)
		return err
	}

	cf.App.APIBaseURL = fmt.Sprintf("%s/%s", cf.App.APIBaseURL, cf.App.Version)
	cf.Swagger.Host = fmt.Sprintf("%s/%s", cf.Swagger.Host, cf.Swagger.Version)
	cf.Validator = validator.New()

	if err := cf.Validator.RegisterValidation("maxString", validateString); err != nil {
		logrus.Error("cannot register maxString Validator config error:", err)
		return err
	}

	if err := cf.Validator.RegisterValidation("intPercent", validateIntPercent); err != nil {
		logrus.Error("cannot register intPercent Validator config error:", err)
		return err
	}

	en := en.New()
	cf.UniversalTranslator = ut.New(en, en)
	enTrans, _ := cf.UniversalTranslator.GetTranslator("en")
	if err := en_translations.RegisterDefaultTranslations(cf.Validator, enTrans); err != nil {
		logrus.Error("cannot add english translator config error:", err)
		return err
	}
	_ = cf.Validator.RegisterTranslation("maxString", enTrans, func(ut ut.Translator) error {
		return ut.Add("maxString", "{0} must have number of characters less than 255", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("maxString", fe.Field())
		return t
	})

	_ = cf.Validator.RegisterTranslation("intPercent", enTrans, func(ut ut.Translator) error {
		return ut.Add("intPercent", "{0} must have value between 0 and 100", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("intPercent", fe.Field())
		return t
	})

	return nil
}

// validateString implements validator.Func for max string by rune
func validateString(fl validator.FieldLevel) bool {
	if lengthOfString := utf8.RuneCountInString(fl.Field().String()); lengthOfString > 255 {
		return false
	}

	return true
}

// validateIntPercent implements validator.Func for int percent 0 to 100
func validateIntPercent(fl validator.FieldLevel) bool {
	if percentInInt := fl.Field().Int(); percentInInt < 0 || percentInInt > 100 {
		return false
	}

	return true
}
