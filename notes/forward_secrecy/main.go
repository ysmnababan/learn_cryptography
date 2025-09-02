package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	mode = flag.String("mode", "server", "server or client")
	// addr = flag.String("addr", "127.0.0.1:4444", "listen/dial address")
	// safe = flag.Bool("safe", false, "server: enable bounds checking (the fix)")
)

func main() {
	flag.Parse()

	switch *mode {
	case "server":
		runServer(":443")
	case "client":
	default:

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

func tlsVersionToString(v uint16) string {
	switch v {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return fmt.Sprintf("Unknown (0x%x)", v)
	}
}

func cipherSuiteToString(cs uint16) string {
	switch cs {
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:
		return "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA"
	case tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA:
		return "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA"
	case tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:
		return "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA"
	case tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA:
		return "TLS_ECDHE_RSA_WITH_RC4_128_SHA"
	case tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA:
		return "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA"
	case tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:
		return "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA"
	case tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:
		return "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA"
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256:
		return "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256"
	case tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256:
		return "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256"
	case tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:
		return "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:
		return "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	case tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:
		return "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
	case tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:
		return "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
	case tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256:
		return "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256"
	case tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256:
		return "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"
	// TLS 1.3 cipher suites.
	case 0x1301:
		return "TLS_AES_128_GCM_SHA256"
	case 0x1302:
		return "TLS_AES_256_GCM_SHA384"
	case 0x1303:
		return "TLS_CHACHA20_POLY1305_SHA256"
	default:
		return fmt.Sprintf("Unknown (0x%x)", cs)
	}
}
