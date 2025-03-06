package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dev-roberto-sousa/nuvemhub/internal/api/handlers"
	"github.com/stretchr/testify/assert"
)

func TestHomeRoute(t *testing.T) {
	router := handlers.SetupRouter()

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"message":"Bem-vindo ao NuvemHub!"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
