package main

import (
	"demo/ReverseProxy/cmd/ReverseProxy/config"
	"demo/ReverseProxy/cmd/ReverseProxy/network"
	"fmt"
	log "github.com/alecthomas/log4go"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync/atomic"
)

type ProxyIpAddress struct {
	target *url.URL
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

func NewMultipleHostsReverseProxy(proxys []*ProxyIpAddress) *httputil.ReverseProxy {
	directorHanlder := func(req *http.Request) {
		proxy := proxys[rand.Int()%len(proxys)]
		atomic.AddUint64(&proxy.Use, 1)
		log.Info("|%-20s | Use:%d", proxy.target.Host, proxy.Use)

		req.URL.Scheme = proxy.target.Scheme
		req.URL.Host = proxy.target.Host
		req.URL.Path = singleJoiningSlash(proxy.target.Path, req.URL.Path)

		targetQuery := proxy.target.RawQuery
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, er error) {
		log.Info(er)
	}

	return &httputil.ReverseProxy{Director: directorHanlder, ErrorHandler: errorHandler}
}

var (
	proxys []*ProxyIpAddress
)

func main() {
	log.LoadConfiguration("./conf/log4go.xml")

	cfg := config.NewConfig()
	if er := cfg.Load("./conf/conf.json"); er != nil {
		return
	}

	for _, h := range cfg.BackgroudURL {
		proxy := &ProxyIpAddress{
			target: &url.URL{
				Scheme: "http",
				Host:   h,
			},
			Use: 0,
		}
		log.Info("backgroud: %s", proxy.target.Host)
		proxys = append(proxys, proxy)
	}

	//go ws.RunEchoMain()

	pk := network.New(1, 1, 1)
	er := pk.Encode([]byte("111111111111111111111111"))
	if er != nil {

	}
	//m := mux.NewRouter()
	//log.Exit(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), m))

	proxyHandler := NewMultipleHostsReverseProxy(proxys)
	log.Exit(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), proxyHandler))
}
