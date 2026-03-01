package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GitLabPayload struct {
	ObjectKind string `json:"object_kind"`
	ProjectID  int    `json:"project_id"`
}

func gitlabWebhookHandler(writing http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writing, "Only POST Requests", http.StatusMethodNotAllowed)
		return
	}

	var payload GitLabPayload

	json.NewDecoder(request.Body).Decode(&payload)

	if payload.ObjectKind == "merge_request" {
		fmt.Printf("Project ID: %d ", payload.ProjectID)
	}
}
