package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

const serverPort = 8080

type GitLabPayload struct {
	ObjectKind string `json:"object_kind"`
	ProjectID  int    `json:"project_id"`

	ObjectAttributes struct {
		IID          int    `json:"iid"`
		State        string `json:"state"`
		SourceBranch string `json:"source_branch"`
		TargetBranch string `json:"target_branch"`
	} `json:"object_attributes"`
}

func gitlabWebhookHandler(writing http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writing, "Only POST Requests", http.StatusMethodNotAllowed)
		return
	}

	var payload GitLabPayload

	json.NewDecoder(request.Body).Decode(&payload)

	if payload.ObjectKind == "merge_request" {
		if payload.ObjectAttributes.State == "opened" {
			fmt.Printf("MR ID: %d \n", payload.ObjectAttributes.IID)
		}
	}
	cmd := exec.Command("sh", "analiz.sh", "./", payload.ObjectAttributes.SourceBranch, payload.ObjectAttributes.TargetBranch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Script çalışırken bir hata oluştu: %s\n", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/webhook", gitlabWebhookHandler)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s\n", err)
		}
	}
}
