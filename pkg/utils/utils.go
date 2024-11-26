package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		fmt.Printf("Request body: %s\n", string(body))
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	} else {
		fmt.Printf("Error reading request body: %v\n", err)
	}
}
