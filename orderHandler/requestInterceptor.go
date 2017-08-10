package interceptor

import "net/http"

type CreateInterceptor struct {
	v int
}

type Interceptor interface {
	Interceptor(w http.ResponseWriter, r *http.Request) (bool, map[string]string)
}

func (i CreateInterceptor)Interceptor(w http.ResponseWriter, r *http.Request) (bool, map[string]string) {

	param := make(map[string]interface{})

	return false, param
}



