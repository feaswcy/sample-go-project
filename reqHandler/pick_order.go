package reqHandler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func handlePKorder(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	resMap := make(map[string]interface{})
	query := r.URL.Query()

	oid := query.Get("oid")
	did := query.Get("did")

	if oid == "" || did == "" {
		resMap["errno"] = 3003
		resMap["errmsg"] = "创建订单缺少参数"

		log := "oid: " + oid + "did : " + oid + "errno: 3003" + "创建订单缺少参数."
		Logger.Println(log)
	}

	oidInt, _ := strconv.Atoi(oid)
	didInt, _ := strconv.Atoi(did)

	res, err := PKOrder(oidInt, didInt, PICKED, PENDING)

	if err != nil {
		resMap["errno"] = 3001
		resMap["errmsg"] = "抢单数据库连接失败"
		log := "oid: " + oid + "did : " + oid + "errno: 3001" + "抢单数据库连接失败," + err.Error()
		Logger.Println(log)
	} else {

		if res > 0 {
			data, err := queryOrdersById(oid)
			if err == nil {
				resMap["errno"] = 0
				resMap["errmsg"] = "抢单成功"
				resMap["data"] = data
				log := "oid: " + oid + "did : " + oid + "抢单成功。"
				Logger.Println(log)
				pushMessage(3, oidInt)            //推送给乘客已接单
				changeGroup(1, 2, didInt, oidInt) //长连接把司机从听单组移到已接单组
			}

		} else {
			resMap["errno"] = 3002
			resMap["errmsg"] = "抢单失败"
			resMap["data"] = "[]"
			log := "oid: " + oid + "did : " + oid + "抢单失败。"
			Logger.Println(log)
		}
	}

	jsonv, _ := json.Marshal(resMap)

	io.WriteString(w, string(jsonv))
}
