package reqHandler

import(

)

func handleCancel(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	requestBody := make(map[string]string)
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), &requestBody)

	resMap := make(map[string]interface{})

	oid := requestBody["oid"]
	status := requestBody["status"]
	role := requestBody["role"]

	statusInt, err := strconv.Atoi(status)
	oidInt, err := strconv.Atoi(oid)
	log.Println(PENDING)
	log.Println(PICKED)

	if err != nil || (statusInt != PENDING && statusInt != PICKED) {

		Logger.Printf("statusInt:%d, pickrd:%d , \r\n" + err.Error(), statusInt, PICKED)
		resMap["errno"] = 5001
		resMap["errmsg"] = "取消状态错误"

		log := "oid: " + oid + "status" + status + " role: " + role + "取消失败" + err.Error()
		Logger.Println(log)

		jsonv, _ := json.Marshal(resMap)
		io.WriteString(w, string(jsonv))
		return
	}

	sql := "UPDATE orders SET status = %d WHERE id = %s and %s = %s and status = %d "

	if role != "driver" && role != "passenger" {

		resMap["errno"] = 5002
		resMap["errmsg"] = "取消角色错误"
		log := "oid: " + oid + "status" + status + " role: " + role + "取消角色错误"
		Logger.Println(log)

		jsonv, _ := json.Marshal(resMap)
		io.WriteString(w, string(jsonv))
		return
	}

	updateStatus := PRE_CANCEL

	if statusInt == PICKED && role == "driver" {
		updateStatus = POST_D_CANCEL
		sql = fmt.Sprintf(sql, updateStatus, oid, "did", requestBody["did"], statusInt)
	} else if statusInt == PICKED && role == "passenger" {
		sql = fmt.Sprintf(sql, updateStatus, oid, "pid", requestBody["pid"], statusInt)
		updateStatus = POST_P_CANCEL
	}
	log.Println(sql)

	cnt, err := update(sql)

	if err != nil {
		resMap["errno"] = 5004
		resMap["errmsg"] = "取消sql语句执行错误"

		log := "oid: " + oid + "status" + status + " role: " + role + "取消sql语句执行错误"
		Logger.Println(log)

	}

	if cnt == 1 {
		resMap["errno"] = 0
		resMap["errmsg"] = "取消成功"
		if role == "driver" {
			pushMessage(3, oidInt)
		} else if role == "passenger" {
			pushMessage(2, oidInt)
		}
	} else {
		resMap["errno"] = 5003
		resMap["errmsg"] = "订单状态不符"
	}

	jsonv, _ := json.Marshal(resMap)
	io.WriteString(w, string(jsonv))
	return
}