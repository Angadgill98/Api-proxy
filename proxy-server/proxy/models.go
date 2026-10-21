package proxy

import (
	"net/http"
	"time"
)


type RegisterAppReq struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}



type UpdateAppReq struct {
    Name    string `json:"name"`
    NewName string `json:"new_name"`
    URL     string `json:"url"`
}



type Application struct {
    ID   int    `json:"id,omitempty"`
    Name string `json:"name"`
    URL  string `json:"url"`
}


type HandleListenersReq struct {
    Operation string `json:"operation"` // add or remove
    AppName   string `json:"app_name"`
}



type ProxyRequest struct {
	DestinationURL string `json:"destination_url"`
	Body           []byte `json:"body"`
}

type ProxyInfo struct {
	ID          int64            `json:"id,omitempty"`
	Route       string            `json:"route"`
	Request     ProxyRequestInfo  `json:"request"`
	Response    ProxyResponseInfo `json:"response"`
	StartedAt   time.Time         `json:"started_at"`
	CompletedAt time.Time         `json:"completed_at"`
	Duration    time.Duration     `json:"duration"`
	// Attempts    int               `json:"attempts"`
	Error       string            `json:"error,omitempty"`
}

type ProxyRequestInfo struct {
	Method        string           `json:"method"`
	URL           string           `json:"url"`
	DestinationURL string          `json:"destination_url"`
	Headers       http.Header      `json:"headers"`
	Cookies       []*http.Cookie   `json:"cookies"`
	Body          []byte           `json:"body"`
}

type ProxyResponseInfo struct {
	StatusCode int            `json:"status_code"`
	Headers    http.Header    `json:"headers"`
	Cookies    []*http.Cookie `json:"cookies"`
	Body       []byte         `json:"body"`
}