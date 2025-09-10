package main

import (
	"crypto/tls"
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	mode = flag.String("mode", "server", "server or client")
)

func runServer(addr string) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/hello", func(c echo.Context) error {
		return c.String(200, "nice !!!")
	})
	s := http.Server{
		Addr:    addr,
		Handler: e,
	}
	err := s.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		panic(err)
	}
}

func runClient(port string) {
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Println("error loading", err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Certificates: []tls.Certificate{cert},
			},
		},
	}
	resp, err := client.Get("https://localhost" + port + "/hello")
	if err != nil {
		log.Fatalf("error get:%v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	log.Printf("Response body: %s\n", body)
}

func main() {
	flag.Parse()

	switch *mode {
	case "server":
		runServer(":8001")
	case "client":
		runClient(":8001")
	default:
		return
	}
}
