package proxy

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"proxy/global"
	"time"
)


func ProxyPipeline(w http.ResponseWriter, req *http.Request) {
	if !ValidateRemoteAddr(req) {
		WriteResponse(w, http.StatusForbidden, false)
		return
	}

	startedAt := time.Now()

	params, err := ParseProxyRequest(req)
	if err != nil {
		log.Println("ProxyPipeline:", err)
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	route, err := GetRoute(params.DestinationURL)
	if err != nil {
		log.Println("ProxyPipeline: invalid destination URL:", err)
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	proxyInfo := NewProxyInfo(req, params, route, startedAt)

	proxy := CreateConfiguredProxy(params.DestinationURL, &proxyInfo)
	proxy.ServeHTTP(w, req)

	CompleteProxyInfo(&proxyInfo)

	if err := SaveProxyInfo(global.Db, proxyInfo); err != nil {
		log.Println("ProxyPipeline: failed to save proxy info:", err)
	}

	log.Printf("Proxy info: %+v\n", proxyInfo)

	WriteResponse(w, http.StatusOK, true, proxyInfo)
}

func ParseProxyRequest(req *http.Request) (ProxyRequest, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return ProxyRequest{}, fmt.Errorf("failed to read request body: %w", err)
	}
	defer req.Body.Close()

	var params ProxyRequest
	if err := json.Unmarshal(body, &params); err != nil {
		return ProxyRequest{}, fmt.Errorf("invalid JSON: %w", err)
	}

	req.Body = io.NopCloser(bytes.NewReader(body))

	return params, nil
}

func NewProxyInfo(req *http.Request, params ProxyRequest, route string, startedAt time.Time) ProxyInfo {
	requestInfo := GetRequestInfo(req)
	requestInfo.DestinationURL = params.DestinationURL

	return ProxyInfo{
		Route:     route,
		Request:   requestInfo,
		StartedAt: startedAt,
	}
}

func CreateConfiguredProxy(destinationURL string, info *ProxyInfo) *httputil.ReverseProxy {
	proxy := CreateProxy(destinationURL)

	proxy.ModifyResponse = func(res *http.Response) error {
		responseInfo, err := GetResponseInfo(res)
		if err != nil {
			return err
		}

		info.Response = responseInfo
		return nil
	}

	proxy.ErrorHandler = func(rw http.ResponseWriter, r *http.Request, err error) {
		info.Error = err.Error()
		log.Println("ProxyPipeline: proxy error:", err)
		http.Error(rw, "Bad Gateway", http.StatusBadGateway)
	}

	return proxy
}

func CompleteProxyInfo(info *ProxyInfo) {
	info.CompletedAt = time.Now()
	info.Duration = info.CompletedAt.Sub(info.StartedAt)
}

func SaveProxyInfo(db *sql.DB, info ProxyInfo) error {
	_, err := UpsertProxyInfo(db, info)
	return err
}



func GetRemoteAddr(req *http.Request) string {
	return req.RemoteAddr
}

func ValidateRemoteAddr(req *http.Request) bool {
	remoteAddr := GetRemoteAddr(req)

	for _, addr := range global.Active_listeners {
		if addr == remoteAddr {
			return true
		}
	}

	return false
}

func GetRoute(destination string) (string, error) {
	parsedURL, err := url.Parse(destination)
	if err != nil {
		return "", err
	}

	if (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") || parsedURL.Host == "" {
		return "", fmt.Errorf("invalid destination URL: %s", destination)
	}

	return parsedURL.Path, nil
}

func GetRequestInfo(req *http.Request) ProxyRequestInfo {
	body, _ := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewReader(body))

	return ProxyRequestInfo{
		Method:  req.Method,
		URL:     req.URL.String(),
		Headers: req.Header.Clone(),
		Cookies: req.Cookies(),
		Body:    body,
	}
}

func GetResponseInfo(res *http.Response) (ProxyResponseInfo, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ProxyResponseInfo{}, err
	}

	res.Body = io.NopCloser(bytes.NewReader(body))

	return ProxyResponseInfo{
		StatusCode: res.StatusCode,
		Headers:    res.Header.Clone(),
		Cookies:    res.Cookies(),
		Body:       body,
	}, nil
}