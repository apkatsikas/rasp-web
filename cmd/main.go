package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/acme/autocert"
)

const (
	timeout      = 5
	idleTimeout  = 120
	httpsPort    = ":443"
	httpPort     = ":80"
	certsDir     = "certs"
	newhellMusic = "newhell.duckdns.org"
)

func main() {
	fmt.Println("TESTING")
	navidrome, err := url.Parse("http://localhost:4533")
	if err != nil {
		panic(fmt.Sprintf("Got error trying to parse navidrome URL %s", err))
	}
	fmt.Println("NewSingleHostReverseProxy")
	navidromeProxy := httputil.NewSingleHostReverseProxy(navidrome)
	naviRouter := chi.NewRouter()
	naviRouter.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		navidromeProxy.ServeHTTP(w, r)
	})

	fmt.Println("http ListenAndServe")
	// Redirect HTTP traffic to https
	go http.ListenAndServe(httpPort, http.HandlerFunc(redirect))

	// Setup certs
	fmt.Println("certs")
	m := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(
			newhellMusic,
			// maybe newhell.link here?
		),
		Cache: autocert.DirCache(certsDir),
	}

	fmt.Println("server")
	// Setup server
	server := &http.Server{
		ReadTimeout:  timeout * time.Second,
		WriteTimeout: timeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
		Addr:         httpsPort,
		TLSConfig:    m.TLSConfig(),
		Handler:      naviRouter,
	}

	fmt.Println("GONNA SERVE")
	// Serve
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		panic(fmt.Sprintf("Got error trying to serve %s", err))
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		fmt.Sprintf("https://%v%v", req.Host, req.URL.String()),
		http.StatusMovedPermanently)
}
