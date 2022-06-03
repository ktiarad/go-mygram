package params

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status         int         `json:"status"`
	Error          string      `json:"error,omitempty"`
	Payload        interface{} `json:"payload,omitempty"`
	Message        string      `json:"message,omitempty"`
	AdditionalInfo interface{} `json:"additional_info,omitempty"`
}

func WriteJsonResponse(rw http.ResponseWriter, payload *Response) {
	rw.WriteHeader(payload.Status)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(payload)
}
