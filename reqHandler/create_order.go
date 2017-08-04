package reqHandler

import (
	
)

func handleCreate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	resMap := make(map[string]interface{})
	// resMap["data"] := make(map[string]string)


	requestBody := make(map[string]string)
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), &requestBody)

	pid := requestBody["pid"]
	from_addr := requestBody["from_addr"]
	to_addr := requestBody["to_addr"]
	//lng := query.Get("lng")
	//lat := query.Get("lat")
	//tips := query.Get("tips")


	if pid == "" || from_addr == "" {
		resMap["errno"] = 4001
		resMap["errmsg"] = "创建订单缺少参数"

		log := "pid: " + pid + "from_addr : " + from_addr + "to_addr" + to_addr + "erro: 4001, 创建订单缺少参数"
		Logger.Println(log)

		jsonv, _ := json.Marshal(resMap)
		io.WriteString(w, string(jsonv))
		return
	}

	pidInt, _ := strconv.Atoi(pid)
	oid, err := insert(pidInt, from_addr, to_addr)
	data := make(map[string]string)

	if err != nil {

		log := "pid: " + pid + "from_addr : " + from_addr + "to_addr" + to_addr + "erro: 4002, 注入订单失败," + err.Error()
		Logger.Println(log)

		resMap["errno"] = 4002
		resMap["errmsg"] = "创建订单失败"

	}

	resMap["errno"] = 0
	resMap["errmsg"] = "创建订单成功"
	data["oid"] = strconv.Itoa(int(oid))
	resMap["data"] = data


	//log := "pid: " + pid + "from_addr : " + from_addr + "to_addr" + to_addr + " oid: " + oid +  "创建订单成功"
	Logger.Println("")

	jsonv, _ := json.Marshal(resMap)

	pushMessage(1, int(oid))
	io.WriteString(w, string(jsonv))
	return
}
