package helpers

import (
	"encoding/json"
	"net/http"
)

const (
	SYSTEM_ERROR = "system error"
)

func HttpResponse(w http.ResponseWriter, out interface{}, statusCode int) {
	w.WriteHeader(statusCode)

	result, err := json.Marshal(out)
	if err == nil {
		w.Write(result)
	}
}

func HttpResponse2(w http.ResponseWriter, statusCode int, out ...interface{}) {
	w.WriteHeader(statusCode)

	if len(out) == 0 {
		return
	}

	result, err := json.Marshal(out[0])
	if err == nil {
		w.Write(result)
	}
}
