package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rakyll/statik/fs"

	_ "github.com/rof20004/brasil-servicos/statik"
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

	// ---------
	// API
	// ---------
	api := echo.New()
	api.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello API Server")
	})
	hosts["api.localhost:3010"] = &Host{api}

	// ---------
	// Site
	// ---------
	site := echo.New()
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	site.GET("/*", echo.WrapHandler(http.StripPrefix("/", http.FileServer(statikFS))))
	hosts["site.localhost:3010"] = &Host{site}

	// http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	// http.ListenAndServe(":3050", nil)

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
