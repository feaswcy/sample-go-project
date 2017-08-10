package service

import (
	"net/http"
	"reqHandler"
	"log"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/websocket"
)

//注册服务接口
func () {
	http.HandleFunc("/listen", handlelisten)
}

//建立长连接，服务端一直查询orderstatus，有改变push到对应的socket
func handlelisten(w http.ResponseWriter, r *http.Request){
	//一个长连接要占用一个端口,双方应该在不同的端口，因此要生成一个标志，来记录双方的关系
}

func pushOrder(){
	
}

func checkOrderStarus(){

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  10240,
	WriteBufferSize: 10240,
}

func handleSocket(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	query := r.URL.Query()

	groupType, _ := strconv.Atoi(query.Get("type"))
	id, _ := strconv.Atoi(query.Get("id"))
	Logger.Println(groupType)
	Logger.Println(id)
	if true {
		r.Header["Origin"] = []string{}
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			Logger.Println(err.Error())
			return
		}

		if err == nil {
			if groupType == 1 {
				Logger.Println("handleSocket err, group type: type1")
				ListeningMap[id] = conn

			} else if groupType == 2 {
				Logger.Println("handleSocket err, group type: type2")
				DriverMap[id] = conn
			} else if groupType == 3 {
				Logger.Println("handleSocket err, group type: type3")
				log.Println("passenger connected")
				PassengerMap[id] = conn
			}
		}

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				Logger.Println(err.Error())
				return
			}
			if err := conn.WriteMessage(messageType, p); err != nil {
				Logger.Println(err.Error())
				return
			}
		}
	}

}


//link_service.go
var ListeningMap = make(map[int]*websocket.Conn)
var DriverMap = make(map[int]*websocket.Conn)
var PassengerMap = make(map[int]*websocket.Conn)
func pushMessage(groupType int, oid int) {
	log.Println("start push message")
	log.Println(groupType)
	log.Println(oid)
	if groupType == 1 {
		for k, v := range ListeningMap {
			log.Println(k)
			if v != nil {
				v.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", oid)))
			}
		}
	} else if groupType == 2 {
		driverSocket := DriverMap[oid]
		if driverSocket != nil {
			driverSocket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", oid)))
		}
	} else if groupType == 3 {
		passengerSocket:= PassengerMap[oid]
		if passengerSocket != nil {
			log.Println("push message to passenger")
			log.Println(oid)
			passengerSocket.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", oid)))
		}
	}
}
func changeGroup(originGroup int, targetGroup int, originKey int, targetKey int) {
	if originGroup == 1 {
		if targetGroup == 2 {
			conn := ListeningMap[originKey]
			delete(ListeningMap, originKey)
			DriverMap[targetKey] = conn
		}
	}
	if originGroup == 2 {
		if targetGroup == 1 {
			conn := DriverMap[originKey]
			delete(DriverMap, originKey)
			ListeningMap[targetKey] = conn
		}
	}
}