package reqHandler

import (
	"encoding/json"
	"io"
	"net/http"
)

func response(w http.ResponseWriter, resMap map[string]interface{}) {
	jsonv, _ := json.Marshal(resMap)

	io.WriteString(w, string(jsonv))
}
