package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	mode = flag.String("mode", "server", "server or client")
	safe = flag.Bool("safe", true, "server: enable bounds checking (the fix)")
)

func main() {
	flag.Parse()

	switch *mode {
	case "server":
		runServer(":443")
	case "client":
		setupClient(":443", *safe)
	default:
		return
	}
}

func runServer(addr string) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/hello", func(c echo.Context) error {
		tls_version := fmt.Sprintf("<p>%s : %s</p>", cipherSuiteToString(c.Request().TLS.CipherSuite), tlsVersionToString(c.Request().TLS.Version))
		return c.HTML(http.StatusOK, `
<h1>Welcome to Echo!</h1>
<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
`+tls_version)
	})
	s := http.Server{
		Addr:    addr,
		Handler: e, // set Echo as handler
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS13,
			CipherSuites: []uint16{
				// DONT USE THIS, RC4 BROKEN, 3DES WEAK
				// tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
				// tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
				// tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,

				// DONT USE THIS, CBC MODES VULNERABLE TO PADDING-ORACLE ATTACKS
				// tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				// tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				// tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				// tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				// tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
				// tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,

				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
			},
		},
	}
	// if err := s.ListenAndServeTLS("server.pem", "server.key"); err != http.ErrServerClosed {
	if err := s.ListenAndServeTLS("server_ec.pem", "server_ec.key"); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}

func setupClient(addr string, isFS bool) {
	conf := tls.Config{
		InsecureSkipVerify: true, // this for testing only, delete this for production
	}
	if isFS {
		conf.CipherSuites = []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		}
	} else {
		conf.CipherSuites = []uint16{
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		}
		conf.MaxVersion = tls.VersionTLS12
	}
	conn, err := tls.Dial("tcp", "127.0.0.1"+addr, &conf)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	req := "GET /hello HTTP/1.1\r\n" +
		"Host: localhost\r\n" +
		"Connection: close\r\n" +
		"\r\n"
	n, err := conn.Write([]byte(req))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 400)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
	}

	println(string(buf[:n]))
}
