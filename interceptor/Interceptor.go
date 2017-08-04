package interceptor

import "net/http"
type Interceptor interface {
	Interceptor(w http.ResponseWriter, r *http.Request) (bool, map[string]string)
}
