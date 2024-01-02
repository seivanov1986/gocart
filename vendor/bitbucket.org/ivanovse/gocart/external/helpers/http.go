package helpers

import (
	"encoding/json"
	"net/http"
)

const (
	SYSTEM_ERROR = "system error"
)

func HttpResponse(w http.ResponseWriter, out interface{}, statusCode int) {
	result, err := json.Marshal(out)
	if err != nil {
		w.Write([]byte(SYSTEM_ERROR + ": " + err.Error()))
		return
	}
	w.WriteHeader(statusCode)

	w.Write(result)
}
