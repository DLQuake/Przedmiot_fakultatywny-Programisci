package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func Decode(r *http.Request, val interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&val); err != nil {
		return NewRequestError(err, http.StatusBadRequest)
	}
	return nil
}