package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func Response (w http.ResponseWriter, response interface{}, status int) {
	data, err := json.Marshal(response)
	if err != nil {
		log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		log.Println("error writing result", err)
		return err
	}
	return nil
}

func Request (r *http.Request, val interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&val)
	if err != nil {
		return err
	}
	return nil
}