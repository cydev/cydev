package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bradfitz/http2"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var (
	addr = flag.String("addr", "localhost:1234", "Addr to listen")
	cert = flag.String("cert", "", "TLS cert")
	key  = flag.String("key", "", "TLS key")
)

func static(c *echo.Context) error {
	return nil
}

func main() {
	flag.Parse()
	e := echo.New()
	e.Use(mw.Recover())
	e.Use(mw.Logger())
	e.Index("static/test.html")
	e.Static("/static", "static")

	var srv http.Server
	srv.Addr = *addr
	srv.Handler = e
	http2.ConfigureServer(&srv, nil)
	if len(*key) != 0 && len(*cert) != 0 {
		log.Fatal(srv.ListenAndServeTLS(*cert, *key))
	} else {
		log.Fatal(srv.ListenAndServe())
	}
}
