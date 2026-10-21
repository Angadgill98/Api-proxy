package proxy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proxy/global"
)

var db=global.Db

func SetUpServerRoutes(){
	http.HandleFunc("/",ProxyPipeline)

	http.HandleFunc("/register-app",RegisterApphandler)

	http.HandleFunc("/update-app",UpdateApphandler)

	http.HandleFunc("/listeners",HandleListeners)

	http.ListenAndServe(":9090",nil)
}





func RegisterApphandler(w http.ResponseWriter, req *http.Request) {
    var params RegisterAppReq

    err := json.NewDecoder(req.Body).Decode(&params)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    fmt.Println("App Name:", params.Name)
    fmt.Println("App URL:", params.URL)

	app:=Application{
		Name: params.Name,
		URL: params.URL,
	}

	app.InsertApp(db,app.Name,app.URL)
}

func UpdateApphandler(w http.ResponseWriter, req *http.Request) {
    var params UpdateAppReq

    err := json.NewDecoder(req.Body).Decode(&params)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }


    fmt.Println("Current Name:", params.Name)
    fmt.Println("New Name:", params.NewName)
    fmt.Println("New URL:", params.URL)

	app:=Application{
		Name: params.NewName,
		URL: params.URL,
	}

	app.UpdateApp(db,params.Name,app.Name,app.URL)
}

func HandleListeners(w http.ResponseWriter, req *http.Request) {
    var data HandleListenersReq

    if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    app, err := GetAppByName(db, data.AppName)
    if err != nil {
        http.Error(w, "Failed to get application", http.StatusInternalServerError)
        return
    }
    if app == nil {
        http.Error(w, "Application not found", http.StatusNotFound)
        return
    }

    switch data.Operation {
    case "add":
        global.Active_listeners[app.URL] = app.URL

    case "remove":
        delete(global.Active_listeners, app.URL)

    default:
        http.Error(w, "Invalid operation", http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Listener " + data.Operation + " successful",
    })
}