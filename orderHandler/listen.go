package poll

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func handleListen(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")

	resMap := make(map[string]interface{})

	//query := r.URL.Query()

	//did := query.Get("did")

	sql := fmt.Sprintf("SELECT id, pid, from_addr, to_addr, price, status, create_time from orders where status = %d order by create_time desc", PENDING)

	res, err := DbQuery(sql)

	if err != nil {

		resMap["errno"] = 1003
		resMap["errmsg"] = "查询链接数据库异常"
		resMap["data"] = err
	}

	resMap["errno"] = 0
	resMap["errmsg"] = "听单成功"

	if res == nil {
		resMap["data"] = "[]"
	} else {
		resMap["data"] = res
	}

	jsonv, _ := json.Marshal(resMap)

	io.WriteString(w, string(jsonv))
	return
}
