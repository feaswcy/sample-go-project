package main

import (
	"fmt"
	"net/http"

	"github.com/widuu/goini"
)

const machineID = 1

var token map[string]map[string]string
var conf = goini.SetConfig("../../conf/app.conf")
var Logger = registerLog()
var oIDGener = uniqIDGener()

func main() {

	//log.SetFlags(log.Lshortfile | log.LstdFlags)

	Logger.Println(conf.GetValue("database", "dbName"))
	fmt.Println(conf.GetValue("database", "dbName"))

	token := make(map[string]map[string]string)

	token["driver"] = make(map[string]string)
	token["passenger"] = make(map[string]string)

	http.HandleFunc("/socket", handleSocket)
	http.HandleFunc("/listen", handleListen)
	http.HandleFunc("/orderList", handleOrderList)
	http.HandleFunc("/cancel", handleCancel)
	http.HandleFunc("/finished", handleFinished)
	http.HandleFunc("/pkorder", handlePKorder)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/create", handleCreate)
	http.HandleFunc("/order/detail", handleOrderDetail)

	err := http.ListenAndServe(conf.GetValue("application", "port"), nil)
	if err != nil {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}
