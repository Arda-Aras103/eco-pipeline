package main

import (
	"fmt"
	"encoding/json"
)

type GitLabPayload struct {
	ObjectKind string `json:"object_kind"`
	ProjectID int `json:"project_id"`
}

