/* Package provider response content render */
package util

import (
	"encoding/json"
	"net/http"
)
//ResponseWithError: when system occur error,return http error code and message
func ResponseWithError(w http.ResponseWriter, status int, message string) {
	responseWithJson(w, status, map[string]string{"error": message,})
}
//ResponseSuccess: when operate successful return http status 200 and payload
func ResponseSuccess(w http.ResponseWriter, payload interface{}) {
	responseWithJson(w, http.StatusOK, payload)
}
//responseWithJson: return payload format json
func responseWithJson(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
