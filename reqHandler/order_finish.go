package reqHandler

import (
	"net/http"
	"strconv"
)

func handleFinished(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	resMap := make(map[string]interface{})
	query := r.URL.Query()

	oid := query.Get("oid")
	did := query.Get("did")

	if oid == "" || did == "" {
		resMap["errno"] = 6003
		resMap["errmsg"] = "用户和订单号为空"
		log := "完成订单失败, 用户和订单号为空"
		Logger.Println(log)
		response(w, resMap)
		return
	}

	oidInt, _ := strconv.Atoi(oid)
	didInt, _ := strconv.Atoi(did)

	res, err := finishedOrder(oidInt, didInt, FINISHED, PICKED)

	if err != nil {
		resMap["errno"] = 6001
		resMap["errmsg"] = "完成订单sql操作失败"
		log := "oid: " + oid + "did: " + did + " 错误代码: " + "6001 " + "完成订单sql操作失败:" + err.Error()
		Logger.Println(log)
	} else {

		if res > 0 {
			resMap["errno"] = 0
			resMap["errmsg"] = "订单已完成"
			log := "oid: " + oid + "did: " + did + " 订单完成."
			Logger.Println(log)
			pushMessage(3, oidInt)
			changeGroup(2, 1, oidInt, didInt) //长连接把司机从听单组移到已接单组
		} else {
			resMap["errno"] = 6002
			resMap["errmsg"] = "未影响订单"
			log := "oid: " + oid + "did: " + did + " 错误代码: " + "6002 " + "完成订单失败"
			Logger.Println(log)
		}
	}

	response(w, resMap)
	return
}
