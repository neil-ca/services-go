package helpers

import (
	"context"
	"encoding/json"
	"net/http"
)

// DecodeEncryptRequest fills struct from JSON details of request
func DecodeEncryptRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request EncryptRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
