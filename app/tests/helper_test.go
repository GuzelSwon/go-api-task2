package tests

import (
	"app/src"
	"app/src/logger"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	testUrl                       = "https://example.com"
	jsonBody                      = TestStruct{"test_value"}
	stringBody                    = "fixed"
	respStatusOk                  = http.StatusOK
	respStatusInternalServerError = http.StatusInternalServerError
)

type TestStruct struct {
	Value string `json:"value"`
}

func getTestData() []*src.Entity {
	var entities []*src.Entity

	jsonFile, _ := os.Open("./testfiles/entities.json")
	byteValue, _ := io.ReadAll(jsonFile)
	jsonErr := json.Unmarshal(byteValue, &entities)

	if jsonErr != nil {
		logger.Fatal(jsonErr.Error(), logrus.Fields{logger.LoggerCategory: logger.LoggerCategoryHTTP})
	}

	return entities
}

func getResponse(server *httptest.Server) *http.Response {
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, _ := client.Do(req)
	return resp
}

func TestFetchEntities(t *testing.T) {
	expected := getTestData()
	body, _ := json.Marshal(expected)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(respStatusOk)
		w.Write(body)
	}))
	defer svr.Close()
	c := src.NewClient(svr.URL)
	actual := c.FetchEntities()
	assert.Equal(t, actual, expected)
}

func TestFetchNothing(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(respStatusOk)
		w.Write(nil)
	}))
	defer svr.Close()
	c := src.NewClient(svr.URL)
	ent := c.FetchEntities()
	assert.Nil(t, ent)
}

func TestCreateRequestWithCorrectMethod(t *testing.T) {
	expectedReq := src.CreateRequest(testUrl, http.MethodGet)
	assert.Equal(t, "GET", expectedReq.Method)
	assert.Equal(t, testUrl, expectedReq.URL.String())
}

func TestCreateRequestWithIncorrectMethod(t *testing.T) {
	expectedReq := src.CreateRequest(testUrl, "incorrectMethod")
	assert.Nil(t, expectedReq)
	// TODO: check Log message
}

func TestSendRequestWithResponse200(t *testing.T) {
	body, _ := json.Marshal(jsonBody)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(respStatusOk)
		w.Write(body)
	}))

	req, _ := http.NewRequest(http.MethodGet, svr.URL, nil)
	expectedResp := src.SendRequest(req)
	assert.Equal(t, 200, expectedResp.StatusCode)
}

func TestSendRequestWithResponse500(t *testing.T) {
	body, _ := json.Marshal(jsonBody)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(respStatusInternalServerError)
		w.Write(body)
	}))

	req, _ := http.NewRequest(http.MethodGet, svr.URL, nil)
	expectedResp := src.SendRequest(req)
	assert.Equal(t, 500, expectedResp.StatusCode)
}

func TestReadResponseBodyWithJsonBody(t *testing.T) {
	body, _ := json.Marshal(jsonBody)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(respStatusOk)
		w.Write(body)
	}))
	resp := getResponse(svr)
	expectedBody := src.ReadResponseBody(resp)
	assert.Equal(t, expectedBody, body)
}

func TestReadResponseBodyWithStringBody(t *testing.T) {
	body, _ := json.Marshal(stringBody)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(respStatusOk)
		w.Write(body)
	}))
	resp := getResponse(svr)

	expectedBody := src.ReadResponseBody(resp)
	assert.Equal(t, expectedBody, body)
}

func TestReadResponseBodyWithEmptyBody(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(respStatusOk)
		w.Write(nil)
	}))
	resp := getResponse(svr)

	expectedBody := src.ReadResponseBody(resp)
	assert.Empty(t, expectedBody)
}

func TestParseResponseBodyWithEntity(t *testing.T) {
	var entities []*src.Entity
	actualEntities := []src.Entity{{1, "someSlug", "someUrl", "someTitle", "someContent",
		"someImage", "someThumbnail", "someStatus", "someCategory",
		"somePublishedAt", "someUpdatedAt", 2}}
	entityByte, _ := json.Marshal(actualEntities)
	expectedEntities := src.ParseResponseBody(entityByte, entities)
	assert.Equal(t, expectedEntities[0], &actualEntities[0])
}

func TestParseResponseBodyEmptyBody(t *testing.T) {
	var entities []*src.Entity

	emptyByte := []byte{}
	parsedEntity := src.ParseResponseBody(emptyByte, entities)
	assert.Nil(t, parsedEntity)
}

func TestParseResponseBodyWithRandomByte(t *testing.T) {
	var entities []*src.Entity

	randomByte := []byte{65}
	parsedEntity := src.ParseResponseBody(randomByte, entities)
	assert.Nil(t, parsedEntity)
}
