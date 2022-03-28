package context

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/core/sql"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

const (
	pathKey            = "path"
	compositeFormDepth = 3
	// UserKey user key
	UserKey = "user"
	// LangKey lang key
	LangKey = "lang"
	// PostgreDatabaseKey database `postgre` key
	PostgreDatabaseKey = "postgre_database"
	// MysqlDatabaseKey database `mysql` key
	MysqlDatabaseKey = "mysql_database"
	// ParametersKey parameters key
	ParametersKey = "parameters"
)

// Context context
type Context struct {
	*fiber.Ctx
}

// New new custom fiber context
func New(c *fiber.Ctx) *Context {
	return &Context{c}
}

// BindValue bind value
func (c *Context) BindValue(i interface{}, validate bool) error {
	switch c.Method() {
	case http.MethodGet:
		_ = c.QueryParser(i)

	default:
		_ = c.BodyParser(i)
	}

	c.PathParser(i, 1)
	c.Locals(ParametersKey, i)
	TrimSpace(i, 1)

	if validate {
		err := c.Validate(i)
		if err != nil {
			return err
		}
	}
	return nil
}

// PathParser parse path param
func (c *Context) PathParser(i interface{}, depth int) {
	formValue := reflect.ValueOf(i)
	if formValue.Kind() == reflect.Ptr {
		formValue = formValue.Elem()
	}
	t := reflect.TypeOf(formValue.Interface())
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		paramValue := formValue.FieldByName(fieldName)
		if paramValue.IsValid() {
			if depth < compositeFormDepth && paramValue.Kind() == reflect.Struct {
				depth++
				c.PathParser(paramValue.Addr().Interface(), depth)
			}
			tag := t.Field(i).Tag.Get(pathKey)
			if tag != "" {
				setValue(paramValue, c.Params(tag))
			}
		}
	}
}

func setValue(paramValue reflect.Value, value string) {
	if paramValue.IsValid() && value != "" {
		switch paramValue.Kind() {
		case reflect.Uint:
			number, _ := strconv.ParseUint(value, 10, 32)
			paramValue.SetUint(number)

		case reflect.String:
			paramValue.SetString(value)

		default:
			number, err := strconv.Atoi(value)
			if err != nil {
				paramValue.SetString(value)
			} else {
				paramValue.SetInt(int64(number))
			}
		}
	}
}

// Validate validate
func (c *Context) Validate(i interface{}) error {
	if err := config.CF.Validator.Struct(i); err != nil {
		return config.RR.CustomMessage(err.Error(), err.Error()).WithLocale(c.Ctx)
	}

	return nil
}

/*
func (c *Context) trimspace(i interface{}) {
	e := reflect.ValueOf(i).Elem()
	for i := 0; i < e.NumField(); i++ {
		if e.Type().Field(i).Type.Kind() != reflect.String {
			continue
		}

		value := e.Field(i).Interface().(string)
		e.Field(i).SetString(strings.TrimSpace(value))
	}
}*/

// Claims jwt claims
type Claims struct {
	jwt.StandardClaims
	Permissions interface{} `json:"permissions"`
}

// GetClaims get user claims
func (c *Context) GetClaims() *Claims {
	user := c.Locals(UserKey).(*jwt.Token)
	return user.Claims.(*Claims)
}

// GetEmployeeID get employee claims
func (c *Context) GetEmployeeID() string {
	token, ok := c.Locals(UserKey).(*jwt.Token)
	if ok {
		cl := token.Claims.(*Claims)
		if cl != nil {
			return c.GetClaims().Subject
		}
	}

	return ""
}

// GetPermissions getpermissions claims
func (c *Context) GetPermissions() string {
	token, ok := c.Locals(UserKey).(*jwt.Token)
	if ok {
		cl := token.Claims.(*Claims)
		if cl != nil {
			return fmt.Sprintf("%v", c.GetClaims().Permissions)
		}
	}

	return ""
}

// GetToken get token
func (c *Context) GetToken() string {
	token, ok := c.Locals(UserKey).(*jwt.Token)
	if ok {

		if token != nil {
			return token.Raw
		}
	}

	return ""
}

// TrimSpace trim space
func TrimSpace(i interface{}, depth int) {
	e := reflect.ValueOf(i).Elem()
	for i := 0; i < e.NumField(); i++ {
		if depth < 3 && e.Type().Field(i).Type.Kind() == reflect.Struct {
			depth++
			TrimSpace(e.Field(i).Addr().Interface(), depth)
		}

		if e.Type().Field(i).Type.Kind() != reflect.String {
			continue
		}

		value := e.Field(i).String()
		e.Field(i).SetString(strings.TrimSpace(value))
	}
}

// GetLevel get level
func (c *Context) GetLevel() string {
	return c.Get("Level")
}

// GetMysqlPrpoDatabase get connection prpo database `mysql`
func (c *Context) GetTestDatabase() *gorm.DB {
	val := c.Locals(MysqlDatabaseKey)
	if val == nil {
		return sql.StartDatabase
	}

	return val.(*gorm.DB)
}

func (c *Context) GetDB() *gorm.DB {
	val := c.Locals(MysqlDatabaseKey)
	if val == nil {
		return sql.StartDatabase
	}

	return val.(*gorm.DB)
}
