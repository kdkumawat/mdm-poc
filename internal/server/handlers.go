package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Policy holds policy details
type Policy struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func PoliciesHandler(w http.ResponseWriter, r *http.Request) {
	policies := []Policy{
		{
			Key:   "LegalNoticeCaption",
			Value: fmt.Sprintf("Hello MDM User - %s", time.Now().Format("15:04:05")),
		},
		{
			Key:   "LegalNoticeText",
			Value: fmt.Sprintf("Welcome to mdm demo %s", time.Now()),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(policies)
}
