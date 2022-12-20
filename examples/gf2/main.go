package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/MortezaRastegarRad/go-admin/adapter/gf2"
	_ "github.com/MortezaRastegarRad/go-admin/modules/db/drivers/mysql"

	"github.com/MortezaRastegarRad/go-admin/engine"
	"github.com/MortezaRastegarRad/go-admin/examples/datamodel"
	"github.com/MortezaRastegarRad/go-admin/modules/config"

	"github.com/MortezaRastegarRad/themes/adminlte"
	"github.com/MortezaRastegarRad/go-admin/modules/language"
	"github.com/MortezaRastegarRad/go-admin/plugins/example"
	"github.com/MortezaRastegarRad/go-admin/template"
	"github.com/MortezaRastegarRad/go-admin/template/chartjs"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	s := g.Server()

	eng := engine.Default()

	cfg := config.Config{
		Env: config.EnvLocal,
		Databases: config.DatabaseList{
			"default": {
				Host:            "127.0.0.1",
				Port:            "3306",
				User:            "root",
				Pwd:             "123456",
				Name:            "godmin",
				MaxIdleConns:    50,
				MaxOpenConns:    150,
				ConnMaxLifetime: time.Hour,
				Driver:          config.DriverMysql,

				//Driver: config.DriverSqlite,
				//File:   "../datamodel/admin.db",
			},
		},
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.CN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}

	template.AddComp(chartjs.NewChart())

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(&cfg).
		AddGenerators(datamodel.Generators).
		AddDisplayFilterXssFilter().
		AddGenerator("user", datamodel.GetUserTable).
		AddPlugins(examplePlugin).
		Use(s); err != nil {
		panic(err)
	}

	s.AddStaticPath("/uploads", "./uploads")

	eng.HTML("GET", "/admin", datamodel.GetContent)

	s.SetPort(9033)
	go s.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()

}
