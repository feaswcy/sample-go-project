package orderHandler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func handleOrderDetail(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	requestBody := make(map[string]string)
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), &requestBody)

	oid := requestBody["oid"]

	respMap := make(map[string]interface{})
	detail, _ := queryOrdersById(oid)
	respMap["errno"] = 0
	respMap["errmsg"] = "获取订单详细信息成功"
	respMap["data"] = detail

	Logger.Println("oid:" + oid + "获取订单详细信息成功")
	jsonv, _ := json.Marshal(respMap)
	io.WriteString(w, string(jsonv))

}
