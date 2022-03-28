package main

import (
	"flag"
	"fmt"

	"github.com/GolfPhumrapee/finance/docs"
	"github.com/GolfPhumrapee/finance/internal/core/config"
	"github.com/GolfPhumrapee/finance/internal/core/sql"
	"github.com/GolfPhumrapee/finance/internal/core/utils"			
	"github.com/GolfPhumrapee/finance/internal/handlers/routes"
)

// @securityDefinitions.apikey ApiKeyAuth
// @name Authorization
// @in ใส่ค่า Bearer เว้นวรรคและตามด้วย TOKEN  ex(Bearer ?????????)
func main() {
	configs := flag.String("config", "configs", "set configs path, default as: 'configs'")
	flag.Parse()

	// Init configuration
	err := config.InitConfig(*configs)
	if err != nil {
		panic(err)
	}

	// =======================================================

	// programatically set swagger info
	docs.SwaggerInfo.Title = config.CF.Swagger.Title
	docs.SwaggerInfo.Description = config.CF.Swagger.Description
	docs.SwaggerInfo.Version = config.CF.Swagger.Version
	docs.SwaggerInfo.Host = config.CF.Swagger.Host
	//=======================================================

	fmt.Println(docs.SwaggerInfo.BasePath)
	fmt.Println(docs.SwaggerInfo.Version)
	fmt.Println(docs.SwaggerInfo.Host)
	// Init return result
	if err := config.InitReturnResult("configs"); err != nil {
		panic(err)
	}

	//=======================================================

	// Init connection mysql
	configuration := sql.Configuration{
		Host:     config.CF.Mysql.Sql.Host,
		Port:     config.CF.Mysql.Sql.Port,
		Username: config.CF.Mysql.Sql.Username,
		Password: config.CF.Mysql.Sql.Password,
	}
	configuration.DatabaseName = config.CF.Mysql.Sql.DatabaseName
	session, err := sql.InitConnectionMysql(configuration)
	if err != nil {
		panic(err)
	}
	sql.StartDatabase = session.Database

	if !config.CF.App.Release {
		sql.SDatabase()
	}
	// Get public key && private key (JWT)
	err = utils.ReadECDSAKey(config.CF.JWT.PrivateKeyPath, config.CF.JWT.PublicKeyPath)
	if err != nil {
		panic(err)
	}

	routes.NewRouter()
}
