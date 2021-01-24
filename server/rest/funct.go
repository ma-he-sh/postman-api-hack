package rest

import (
	"encoding/json"
	"net/http"
)

func APIPayload(w http.ResponseWriter, data interface{}, funcName string, isError bool) {
	msgType := "success"
	if isError {
		msgType = "error"
	}

	response := map[string]interface{}{
		"resp":    funcName,
		"data":    data,
		"type":    msgType,
		"version": RESTVersion(),
		"server":  RESTServerName(),
	}

	jsonRestData(w, response)
}

func jsonRestData(w http.ResponseWriter, data interface{}) {
	senddata, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(senddata)
}
