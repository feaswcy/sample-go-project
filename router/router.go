package router

import (
	"log"
	"net/http"
	"orderHandler"
)

func Router() {
	http.HandleFunc("/orderList", orderHandler.handleOrderList)
	http.HandleFunc("/cancel", orderHandler.handleCancel)
	http.HandleFunc("/finished", orderHandler.handleFinished)
	http.HandleFunc("/pkorder", orderHandler.handlePKorder)
	http.HandleFunc("/create", orderHandler.handleCreate)
	http.HandleFunc("/orderDetail", orderHandler.handleOrderDetail)

	http.HandleFunc("/login", account.handleLogin)
	http.HandleFunc("/register", account.handleRegister)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
