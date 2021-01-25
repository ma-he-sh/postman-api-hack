package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RestRoutes(routes *mux.Router) {
	routes.HandleFunc("/api/status", restStatus).Methods("GET", "POST")
	routes.HandleFunc("/api/fetch", restFetch).Methods("POST")
	routes.HandleFunc("/api/list", restList).Methods("GET")
	routes.HandleFunc("/api/version", restVersion).Methods("GET")
}

// status check
func restStatus(w http.ResponseWriter, r *http.Request) {
	payload := map[string]interface{}{
		"status": "up",
	}
	APIPayload(w, payload, "rest_status", false)
	return
}

// get payload
func restFetch(w http.ResponseWriter, r *http.Request) {
	var codes RequestCodes
	err := json.NewDecoder(r.Body).Decode(&codes)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	reqCodes := GetByCodes(codes)

	APIPayload(w, reqCodes, "rest_fetch", false)
	return
}

// get a list of code list :: for autocomplete cli
func restList(w http.ResponseWriter, r *http.Request) {
	list := GetListOfLang()
	payload := map[string]interface{}{
		"codes": list.Codes,
		"hash":  list.Hash,
	}
	APIPayload(w, payload, "rest_list", false)
	return
}

// get version of payload :: for update
func restVersion(w http.ResponseWriter, r *http.Request) {
	version := GetRestVersion()
	APIPayload(w, version, "rest_version", false)
	return
}
