package proxy

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"proxy/global"
	"time"
)

var db=global.Db

func SetUpServerRoutes(){
	http.HandleFunc("/",ProxyPipeline)

	http.HandleFunc("/register-app",RegisterApphandler)

	http.HandleFunc("/update-app",UpdateApphandler)

	http.HandleFunc("/listeners",HandleListeners)

	http.HandleFunc("/update-req",HandleReqUpdate)

	http.HandleFunc("/execute-req",HandleExecute)

	http.ListenAndServe(":9090",nil)
}



func WriteResponse(w http.ResponseWriter, status int, success bool, data ...any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var responseData any
	if len(data) > 0 {
		responseData = data[0]
	}

	json.NewEncoder(w).Encode(map[string]any{
		"success": success,
		"data":    responseData,
	})
}

func RegisterApphandler(w http.ResponseWriter, req *http.Request) {
	var params RegisterAppReq

	if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
		log.Println("RegisterApphandler: invalid JSON:", err)
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	app := Application{Name: params.Name, URL: params.URL}

	if err := app.InsertApp(db, app.Name, app.URL); err != nil {
		log.Println("RegisterApphandler: failed to insert application:", err)
		WriteResponse(w, http.StatusInternalServerError, false)
		return
	}

	WriteResponse(w, http.StatusCreated, true)
}

func UpdateApphandler(w http.ResponseWriter, req *http.Request) {
	var params UpdateAppReq

	if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
		log.Println("UpdateApphandler: invalid JSON:", err)
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	app := Application{Name: params.NewName, URL: params.URL}

	if err := app.UpdateApp(db, params.Name, app.Name, app.URL); err != nil {
		log.Println("UpdateApphandler: failed to update application:", err)
		WriteResponse(w, http.StatusInternalServerError, false)
		return
	}

	WriteResponse(w, http.StatusOK, true)
}

func HandleListeners(w http.ResponseWriter, req *http.Request) {
	var data HandleListenersReq

	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		log.Println("HandleListeners: invalid JSON:", err)
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	app, err := GetAppByName(db, data.AppName)
	if err != nil {
		log.Println("HandleListeners: failed to get application:", err)
		WriteResponse(w, http.StatusInternalServerError, false)
		return
	}

	if app == nil {
		log.Println("HandleListeners: application not found:", data.AppName)
		WriteResponse(w, http.StatusNotFound, false)
		return
	}

	switch data.Operation {
	case "add":
		global.Active_listeners[app.URL] = app.URL
	case "remove":
		delete(global.Active_listeners, app.URL)
	default:
		log.Println("HandleListeners: invalid operation:", data.Operation)
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	WriteResponse(w, http.StatusOK, true)
}

func HandleReqUpdate(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPut {
		WriteResponse(w, http.StatusMethodNotAllowed, false)
		return
	}
	defer req.Body.Close()

	var proxyInfo ProxyInfo
	if err := json.NewDecoder(req.Body).Decode(&proxyInfo); err != nil {
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	if proxyInfo.ID == 0 {
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	exists, err := ProxyInfoExists(global.Db, proxyInfo.ID)
	if err != nil {
		log.Println("HandleReqUpdate: failed to check proxy info:", err)
		WriteResponse(w, http.StatusInternalServerError, false)
		return
	}

	if !exists {
		WriteResponse(w, http.StatusNotFound, false)
		return
	}

	if _, err := UpsertProxyInfo(global.Db, proxyInfo); err != nil {
		log.Println("HandleReqUpdate: failed to update proxy info:", err)
		WriteResponse(w, http.StatusInternalServerError, false)
		return
	}

	WriteResponse(w, http.StatusOK, true)
}

func HandleExecute(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		WriteResponse(w, http.StatusMethodNotAllowed, false)
		return
	}
	defer req.Body.Close()

	var proxyInfo ProxyInfo
	if err := json.NewDecoder(req.Body).Decode(&proxyInfo); err != nil {
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	if proxyInfo.ID == 0 {
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	exists, err := ProxyInfoExists(global.Db, proxyInfo.ID)
	if err != nil {
		log.Println("HandleExecute: failed to check proxy info:", err)
		WriteResponse(w, http.StatusInternalServerError, false)
		return
	}
	if !exists {
		WriteResponse(w, http.StatusNotFound, false)
		return
	}

	proxyInfo.StartedAt = time.Now()
	proxyInfo.CompletedAt = time.Time{}
	proxyInfo.Duration = 0
	proxyInfo.Error = ""
	proxyInfo.Response = ProxyResponseInfo{}

	outReq, err := http.NewRequestWithContext(
		req.Context(),
		proxyInfo.Request.Method,
		proxyInfo.Request.DestinationURL,
		bytes.NewReader(proxyInfo.Request.Body),
	)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, false)
		return
	}

	outReq.Header = proxyInfo.Request.Headers.Clone()
	for _, cookie := range proxyInfo.Request.Cookies {
		outReq.AddCookie(cookie)
	}

	proxy := CreateConfiguredProxy(proxyInfo.Request.DestinationURL, &proxyInfo)
	proxy.ServeHTTP(w, outReq)

	CompleteProxyInfo(&proxyInfo)

	if _, err := UpsertProxyInfo(global.Db, proxyInfo); err != nil {
		log.Println("HandleExecute: failed to update proxy info:", err)
		return
	}

	WriteResponse(w, http.StatusOK, true, proxyInfo)
}