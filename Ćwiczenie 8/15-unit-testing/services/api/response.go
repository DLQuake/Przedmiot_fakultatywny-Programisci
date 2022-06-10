package api

import (
	"encoding/json"
	"errors"
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

func ResponseError (w http.ResponseWriter, err error) error{
	if webErr, ok := errors.Cause(err).(*Error); ok {
		er := ErrorResponse{
			Error: webErr.Error(),
		}
		if err := Response(w, er, webErr.Status); err != nil {
			return err
		}
		return nil
	}

	er := ErrorResponse{
		Error: http.StatusText(http.StatusInternalServerError),
	}
	if err := Response(w, er, http.StatusInternalServerError); err != nil {
		return err
	}
	return nil
}