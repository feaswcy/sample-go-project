package reqHandler

import(
	"log"
	"net/http"
) 


func handleLogin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	query := r.URL.Query()

	phone := query.Get("phone")
	role := query.Get("role")

	resMap := make(map[string]interface{})

	flag, user, err := isRegister(phone, role)

	if err != nil {
		resMap["errno"] = 2001
		resMap["errmsg"] = "登录查询用户失败"
		log := "phone: " + phone + ", 查询用户是否注册失败," + err.Error()
		Logger.Println(log)
		response(w, resMap)
		return
	}

	if !flag {
		resMap["errno"] = 2002
		resMap["errmsg"] = "用户未注册"
		log := "phone: " + phone + ", 用户未注册"
		Logger.Println(log)
		response(w, resMap)
		return
	}

	token := genToken(phone, role)

	res, err := updateToken(phone, role, token)

	if err != nil || res <= 0 {
		resMap["errno"] = 2003
		resMap["errmsg"] = "更新Token失败"
		log := "phone: " + phone + ", role: " + role + ",更新token失败, " + err.Error()
		Logger.Println(log)
		response(w, resMap)
		return
	} else {
		resMap["errno"] = 0
		resMap["errmsg"] = "update token success"
		resMap["data"] = user
		log := "phone: " + phone + ", role: " + role + ",update token success"
		Logger.Println(log)
	}

	response(w, resMap)
	return
}
