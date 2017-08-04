package main

import(

	"log"

	"fmt"
	"github.com/gorilla/websocket"
)


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











