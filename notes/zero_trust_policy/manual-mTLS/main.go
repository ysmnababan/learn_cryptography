package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	mode = flag.String("mode", "server", "server or client")
)

func printHeader(r *http.Request) {
	log.Print(">>>>>>>>>>>>>>>> Header <<<<<<<<<<<<<<<<")
	// Loop over header names
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			log.Printf("%v:%v", name, value)
		}
	}
}

func printConnState(state *tls.ConnectionState) {
	log.Print(">>>>>>>>>>>>>>>> State <<<<<<<<<<<<<<<<")

	log.Printf("Version: %x", state.Version)
	log.Printf("HandshakeComplete: %t", state.HandshakeComplete)
	log.Printf("DidResume: %t", state.DidResume)
	log.Printf("CipherSuite: %x", state.CipherSuite)
	log.Printf("NegotiatedProtocol: %s", state.NegotiatedProtocol)

	log.Print("Certificate chain:")
	for i, cert := range state.PeerCertificates {
		subject := cert.Subject
		issuer := cert.Issuer
		log.Printf(" %d s:/C=%v/ST=%v/L=%v/O=%v/OU=%v/CN=%s", i, subject.Country, subject.Province, subject.Locality, subject.Organization, subject.OrganizationalUnit, subject.CommonName)
		log.Printf("   i:/C=%v/ST=%v/L=%v/O=%v/OU=%v/CN=%s", issuer.Country, issuer.Province, issuer.Locality, issuer.Organization, issuer.OrganizationalUnit, issuer.CommonName)
	}
}

func runServer(addr string) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/hello", func(c echo.Context) error {
		printHeader(c.Request())
		printConnState(c.Request().TLS)
		return c.String(200, "nice !!!")
	})
	ca, err := os.ReadFile("./certs/ca.crt")
	if err != nil {
		log.Println("error read ca", err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(ca)

	conf := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  certPool,
	}
	s := http.Server{
		Addr:      addr,
		Handler:   e,
		TLSConfig: conf,
	}

	err = s.ListenAndServeTLS("./certs/server.crt", "./certs/server.key")
	if err != nil {
		panic(err)
	}
}

func runClient(port string) {
	ca, err := os.ReadFile("./certs/ca.crt")
	if err != nil {
		log.Println("error read ca", err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(ca)
	cert, err := tls.LoadX509KeyPair("./certs/client.a.crt", "./certs/client.a.key")
	if err != nil {
		log.Println("error loading", err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      certPool,
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
