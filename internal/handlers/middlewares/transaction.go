package middlewares

import (
	"gorm.io/gorm"

	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/GolfPhumrapee/finance/internal/core/sql"
	"github.com/gofiber/fiber/v2"
)

type (
	// Skipper defines a function to skip middleware. Returning true skips processing
	Skipper func(*fiber.Ctx) bool
)

// TransactionMysql to do transaction mysql
func TransactionMysql(skipper Skipper) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		database := &gorm.DB{}
		skip := skipper(c)
		defer func() {
			if r := recover(); r != nil {
				if !skip {
					_ = database.Rollback()
				}

				stackTrace(c, r)
			}
		}()

		if !skip {
			database = sql.StartDatabase.Begin()
			c.Locals(context.MysqlDatabaseKey, database)
			err = c.Next()
			if err != nil {
				_ = database.Rollback()
			} else {
				if database.Commit().Error != nil {
					_ = database.Rollback()
				}
			}
		} else {
			return c.Next()
		}

		return
	}
}
