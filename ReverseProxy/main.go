package main

import (
	"demo/ReverseProxy/ReverseProxy/config"
	"demo/ReverseProxy/ReverseProxy/network"
	"fmt"
	log "github.com/alecthomas/log4go"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync/atomic"
)

type ProxyConfig struct {
	back   *url.URL
	Use    uint64
	Enable bool
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func NewMultipleHostsReverseProxy(proxyConfig []*ProxyConfig) *httputil.ReverseProxy {
	directorHandler := func(req *http.Request) {
		proxy := proxyConfig[rand.Int()%len(proxyConfig)]
		atomic.AddUint64(&proxy.Use, 1)
		log.Info("|%-20s | Use:%d", proxy.back.Host, proxy.Use)

		req.URL.Scheme = proxy.back.Scheme
		req.URL.Host = proxy.back.Host
		req.URL.Path = singleJoiningSlash(proxy.back.Path, req.URL.Path)

		targetQuery := proxy.back.RawQuery
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}

		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "go-levi")
		}
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, er error) {
		log.Info(er)
	}

	return &httputil.ReverseProxy{Director: directorHandler, ErrorHandler: errorHandler}
}

func main() {
	log.LoadConfiguration("conf/log4go.xml")
	cfg := config.NewConfig()
	if err := cfg.Load("conf/conf.json"); err != nil {
		return
	}

	var proxyConfig []*ProxyConfig
	for _, background := range cfg.Background {
		u, err := url.ParseRequestURI(background)
		if err == nil {
			proxy := &ProxyConfig{
				back: u,
				Use:  0,
			}
			log.Info("background: %s", proxy.back.Host)
			proxyConfig = append(proxyConfig, proxy)
		}
	}

	// go ws.RunEchoMain()

	pk := network.New(1, 1, 1)
	err := pk.Encode([]byte("111111111111111111111111"))
	if err != nil {

	}

	// m := mux.NewRouter()
	// log.Exit(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), m))

	proxyHandler := NewMultipleHostsReverseProxy(proxyConfig)
	log.Exit(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), proxyHandler))
}
