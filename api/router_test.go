package api_test

import (
	"bytes"
	"github.com/angel-one/go-example-project/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/angel-one/go-example-project/api"
	"github.com/angel-one/go-example-project/constants"
	"github.com/stretchr/testify/assert"
)

func testAPI(t *testing.T, request *http.Request, expectedStatus int, expectedBody string) {
	router := api.GetRouter()
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	assert.Equal(t, expectedStatus, w.Code)
	assert.Equal(t, expectedBody, w.Body.String())
}

func TestPing(t *testing.T) {
	body, err := json.Marshal(models.FullNameRequest{
		FirstName: "Naruto",
		LastName:  "Rocks",
	})
	assert.NoError(t, err)
	request, err := http.NewRequest(http.MethodPost, constants.FullNameRoute, bytes.NewReader(body))
	assert.NoError(t, err)
	response, err := json.Marshal(models.FullNameResponse{FullName: "Naruto Rocks"})
	assert.NoError(t, err)
	testAPI(t, request, http.StatusOK, string(response))
}
