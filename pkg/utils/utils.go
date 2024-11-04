package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseJsonBody(r *http.Request, x interface{}) error {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return err
		}
	}
	return nil
}
