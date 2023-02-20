package domainhttp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

var basePath = "/v1/todos"

var handler TodoHandler

func init() {
	handler = NewTodoHandler()
}

func TestTodoHealthCheck(t *testing.T) {
	path := "/health"
	req, err := http.NewRequest("GET", basePath+path, nil)
	require.Nil(t, err)

	router := gin.Default()

	router.GET(basePath+path, handler.Health)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	fmt.Printf("%+v\n", rr)
	fmt.Printf("Body %v\n", rr.Body.String())

	require.Equal(t, 200, rr.Code)
}
