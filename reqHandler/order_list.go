package reqHandler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func handleOrderList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	resMap := make(map[string]interface{})

	//query := r.URL.Query()

	requestBody := make(map[string]string)
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), &requestBody)

	log.Println(requestBody)
	pid := requestBody["pid"]
	did := requestBody["did"]
	status := requestBody["staus"]

	log.Println(pid)
	log.Println(did)
	log.Println(status)

	if pid == "" && did == "" {
		resMap["errno"] = 1001
		resMap["errmsg"] = "pid和did均为空"

		log := "查询订单失败, pid和did均为空, errno: 1001."
		Logger.Println(log)

		jsonv, _ := json.Marshal(resMap)
		io.WriteString(w, string(jsonv))
		return
	}

	if pid != "" && did != "" {
		resMap["errno"] = 1002
		resMap["errmsg"] = "pid和did均不为空"

		log := "查询订单失败, pid和did均不为空, errno: 1002."
		Logger.Println(log)

		jsonv, _ := json.Marshal(resMap)
		io.WriteString(w, string(jsonv))
		return
	}

	sqlStr := getSql(pid, did, status)

	res, err := DbQuery(sqlStr)

	if err != nil {

		resMap["errno"] = 1003
		resMap["errmsg"] = "查询链接数据库异常"
		resMap["data"] = err

		log := "查询链接数据库异常, errno: 1003," + err.Error()
		Logger.Println(log)
	}

	resMap["errno"] = 0
	resMap["errmsg"] = "查询成功"
	log := "id: " + pid + did + " 查询成功。"
	Logger.Println(log)

	if res == nil {
		resMap["data"] = "[]"
	} else {
		resMap["data"] = res
	}

	jsonv, _ := json.Marshal(resMap)
	io.WriteString(w, string(jsonv))
	return
}
