package main

import (
	"proxy/db"
	"proxy/global"
	"proxy/proxy"
)


func main(){
	DB:=global.Db
	db.CreateAppTable(DB)
	db.CreateProxyInfoTable(DB)
	proxy.SetUpServerRoutes()
}