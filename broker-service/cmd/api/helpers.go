package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *htpp.Request, data any) error {
	maxBytes := 1048576 // 1 MB

	r.Body = http.MaxBytesReader(w, r.body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	

	err = dec.Decode(&struct{};
}