package main

import (
	"bytes"
	"flag"
	"log"

	"github.com/valyala/fasthttp"
)

const (
	version = "1.1a"
)

var (
	addr    = flag.String("addr", "localhost:1234", "Addr to listen")
	addrTLS = flag.String("addrTLS", "", "Addr to listen TLS")
	cert    = flag.String("cert", "", "TLS cert")
	key     = flag.String("key", "", "TLS key")
)

func NewPathStripper() fasthttp.PathRewriteFunc {
	var staticPrefix = []byte("/static")
	return func(ctx *fasthttp.RequestCtx) []byte {
		path := ctx.Path()
		if bytes.HasPrefix(path, staticPrefix) {
			path = path[len(staticPrefix):]
		}
		return path
	}
}

func main() {
	flag.Parse()

	fs := &fasthttp.FS{
		Root:               "static",
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: false,
		Compress:           true,
		AcceptByteRange:    true,
		PathRewrite:        NewPathStripper(),
	}

	fsHandler := fs.NewRequestHandler()
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch {
		case bytes.HasPrefix(ctx.Path(), []byte("/hath")):
			fasthttp.ServeFile(ctx, "static/hath.html")
		default:
			fsHandler(ctx)
		}
	}

	// Start HTTP server.
	if len(*addr) > 0 {
		log.Printf("Starting HTTP server on %q", *addr)
		go func() {
			if err := fasthttp.ListenAndServe(*addr, requestHandler); err != nil {
				log.Fatalf("error in ListenAndServe: %s", err)
			}
		}()
	}

	// Start HTTPS server.
	if len(*addrTLS) > 0 {
		log.Printf("Starting HTTPS server on %q", *addrTLS)
		go func() {
			if err := fasthttp.ListenAndServeTLS(*addrTLS, *cert, *key, requestHandler); err != nil {
				log.Fatalf("error in ListenAndServeTLS: %s", err)
			}
		}()
	}
	select {}
}
