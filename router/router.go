package router

import (
	"net/http"
)

func Router() {
	http.HandleFunc("/socket", reqhandleSocket)
	http.HandleFunc("/listen", handleListen)
	http.HandleFunc("/orderList", handleOrderList)
	http.HandleFunc("/cancel", handleCancel)
	http.HandleFunc("/finished", handleFinished)
	http.HandleFunc("/pkorder", handlePKorder)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/create", handleCreate)
	http.HandleFunc("/order/detail", handleOrderDetail)
}
