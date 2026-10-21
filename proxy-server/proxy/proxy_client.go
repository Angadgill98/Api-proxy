package proxy

import (
	"net/http/httputil"
	"net/url"
)

func CreateProxy(target string) *httputil.ReverseProxy {
	destination, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(destination)
}