package Utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&data); err != nil {
		response := map[string]string{"message": err.Error()}
		ResponseJSON(w, http.StatusBadRequest, response)
		return err
	}

	return nil
}