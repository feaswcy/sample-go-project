package poll

import (
	"log"
	"net/http"
	"strconv"

	"log"

	"github.com/gorilla/websocket"
	//"github.com/zheng-ji/goSnowFlake"
)

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
