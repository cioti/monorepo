package client_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cioti/monorepo/cms.api/shared"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var id = uuid.MustParse("380f9ef6-d2c7-4c0f-a3c2-3a782111fc1e")

func TestIntegration_CreateProject(t *testing.T) {
	payload := shared.CreateProjectCommand{
		ID:          id,
		Name:        "some name",
		Description: "some description",
	}
	json, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:8080/v1/projects", "application/json", bytes.NewBuffer(json))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestIntegration_AddModel(t *testing.T) {
	payload := shared.AddModelCommand{
		ProjectID:   id,
		ApiID:       "model-api-id2",
		Name:        "model name2",
		Description: "model description2",
	}
	json, _ := json.Marshal(payload)

	resp, err := http.Post(fmt.Sprintf("http://localhost:8080/v1/projects/%s/models", id), "application/json", bytes.NewBuffer(json))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
