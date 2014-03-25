package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
)

var (
	basedir  string
	bindAddr string
	prefix   string
	esPrefix string
	esAddr   string
)

func init() {
	flag.StringVar(&bindAddr, "bind", ":9090", "bind address")
	flag.StringVar(&prefix, "prefix", "/logstash/", "uri prefix")
	flag.StringVar(&basedir, "basedir", "./logstash-1.4.0/vendor", "base dir")
	flag.StringVar(&esPrefix, "esprefix", "/es/", "elasticsearch prefix")
	flag.StringVar(&esAddr, "esaddr", "http://127.0.0.1:9200/", "elasticsearch address")
}

func handleStriped(prefix string, handler http.Handler) {
	n := len(prefix)
	if n > 0 && prefix[n-1] != '/' {
		prefix += "/"
	}
	http.Handle(prefix, http.StripPrefix(prefix, handler))
}

func main() {

	flag.Parse()

	handleStriped(prefix, http.FileServer(http.Dir(basedir)))

	esPrefix = path.Join(prefix, esPrefix)
	esUrl, _ := url.Parse(esAddr)

	handleStriped(esPrefix, httputil.NewSingleHostReverseProxy(esUrl))

	err := http.ListenAndServe(bindAddr, nil)
	if err != nil {
		panic(err)
	}
}
