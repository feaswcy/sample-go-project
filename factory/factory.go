package factory

import (
	"main/interceptor"
	"main/controller"
	"net/http"
	"encoding/json"
	"io"
)

type Factory interface {
	create(i interceptor.Interceptor, c controller.Controller)
}

type HandleFuncFactory struct {

}

func (h *HandleFuncFactory)create(i interceptor.Interceptor, c controller.Controller) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		flag, param := i.Interceptor(w, r)
		if !flag {
			return
		}

		resMap := c.Deal(param)
		jsonv, _ := json.Marshal(resMap)
		io.WriteString(w, string(jsonv))

		return
	}
}