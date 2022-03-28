package routes

import (
	ctx "context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/GolfPhumrapee/finance/internal/core/utils"
	"github.com/GolfPhumrapee/finance/internal/handlers/middlewares"
	"github.com/GolfPhumrapee/finance/internal/pkg/delete"
	deletestatus "github.com/GolfPhumrapee/finance/internal/pkg/deleteStatus"
	"github.com/GolfPhumrapee/finance/internal/pkg/insertData"
	"github.com/GolfPhumrapee/finance/internal/pkg/login"
	"github.com/GolfPhumrapee/finance/internal/pkg/logout"
	"github.com/GolfPhumrapee/finance/internal/pkg/refreshTK"
	"github.com/GolfPhumrapee/finance/internal/pkg/selectGroup"
	upload "github.com/GolfPhumrapee/finance/internal/pkg/upLoad"
	"github.com/GolfPhumrapee/finance/internal/pkg/updateData"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/form3tech-oss/jwt-go"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	jwtware "github.com/gofiber/jwt/v2"

	"github.com/sirupsen/logrus"
)

const (
	// MaximumSize100MB body limit 100 mb.
	MaximumSize100MB = 1024 * 1024 * 100
	// MaximumSize1MB body limit 1 mb.
	MaximumSize1MB = 1024 * 1024 * 1
)

// NewRouter new router
func NewRouter() {
	app := fiber.New(
		fiber.Config{
			IdleTimeout:    5 * time.Second,
			BodyLimit:      MaximumSize100MB,
			ReadBufferSize: MaximumSize1MB,
		},
	)

	app.Use(
		compress.New(),
		requestid.New(),
		cors.New(),
		middlewares.Logger(),
		middlewares.WrapError(),
		middlewares.TransactionMysql(func(c *fiber.Ctx) bool {
			return c.Method() == fiber.MethodGet
		}),
	)

	app.Static("", "./public", fiber.Static{
		Compress: true,
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Use(middlewares.AcceptLanguage())
	if config.CF.Swagger.Enable {
		v1.Get("/swagger/*", swagger.HandlerDefault)
	}

	authJWT := jwtware.New(jwtware.Config{
		Claims:        &context.Claims{},
		SigningMethod: jwt.SigningMethodES256.Name,
		SigningKey:    utils.VerifyKey,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.
				Status(config.RR.Internal.Unauthorized.HTTPStatusCode()).
				JSON(config.RR.Internal.Unauthorized.WithLocale(c))
		},
	})

	//login && Gen Token
	Login := login.NewEndpoint()
	login := v1.Group("login")
	login.Post("/login", Login.Login)
	//refreshToken
	refreshTK := refreshTK.NewEndpoint()
	refresh := v1.Group("test", authJWT)
	refresh.Post("/refreshTK", refreshTK.RETK)
	//logout
	testlogout := logout.NewEndpoint()
	logout := v1.Group("test", authJWT)
	logout.Post("/logout", testlogout.Logout)
	//insert
	insert := insertData.NewEndpoint()
	I := v1.Group("Data", authJWT)
	I.Post("/InsertData", insert.InsertData)
	//select group
	selectG := selectGroup.NewEndpoint()
	S := v1.Group("Data", authJWT)
	S.Post("/SelectGroup", selectG.SelectGroup)
	//update
	update := updateData.NewEndpoint()
	U := v1.Group("Data", authJWT)
	U.Post("/UpdateData", update.Update)
	//uploadfile
	Upload := upload.NewEndpoint()
	UL := v1.Group("upload", authJWT)
	UL.Post("/Upload", Upload.UpLoads)
	//delete
	delete := delete.NewEndpoint()
	D := v1.Group("Data", authJWT)
	D.Post("/DeleteData", delete.Delete)
	//delete status
	deleteS := deletestatus.NewEndpoint()
	DS := v1.Group("Data", authJWT)
	DS.Post("/DeleteStatus", deleteS.Deletestatus)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		_, cancel := ctx.WithTimeout(ctx.Background(), 5*time.Second)
		defer cancel()

		logrus.Info("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	logrus.Infof("Start server on port: %d ...", config.CF.App.Port)
	err := app.Listen(fmt.Sprintf(":%d", config.CF.App.Port))
	if err != nil {
		logrus.Panic(err)
	}
}
