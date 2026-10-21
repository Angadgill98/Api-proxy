package proxy


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