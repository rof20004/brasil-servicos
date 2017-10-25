package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rakyll/statik/fs"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/rof20004/brasil-servicos/statik"

	"github.com/rof20004/brasil-servicos/api/routes"
	"github.com/rof20004/brasil-servicos/database"
)

type (
	// Host - all subdomains
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	// Hosts
	hosts := map[string]*Host{}

	// Start databse and migrations
	database.InitDatabase()
	database.InitMigrations()

	// Start routes
	routes.InitRoutes()

	// ---------
	// API host
	// ---------
	hosts["api.localhost:3010"] = &Host{routes.API}

	// ---------
	// Site host
	// ---------
	site := echo.New()
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	site.GET("/*", echo.WrapHandler(http.StripPrefix("/", http.FileServer(statikFS))))
	hosts["site.localhost:3010"] = &Host{site}

	// ---------
	// Server
	// ---------
	e := echo.New()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})

	e.Logger.Fatal(e.Start(":3010"))
}
