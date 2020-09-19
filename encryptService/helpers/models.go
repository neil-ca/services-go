package helpers

// EncryptRequest structures request coming from client
type EncryptRequest struct {
	Text string `json:"text"`
	Key string `json:"key"`
}

// EncryptResponse structures response going to the client
type EncryptResponse struct {
	Message string `json:"message"`
	Err string `json:"error"`
}

// DecryptRequest structures request coming from client
type DecryptRequest struct {
	Text string `json:"message"`
	Key string `json:"key"`
}

// DecryptResponse structures response going to the client
type DecryptResponse struct {
	Message string `json:"text"`
	Err string `json:"error"`
}