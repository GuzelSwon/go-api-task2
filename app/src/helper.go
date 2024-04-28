package src

import (
	"app/src/logger"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}
func (c Client) FetchEntities() []*Entity {
	var entities []*Entity

	req := CreateRequest(c.url, http.MethodGet)
	resp := SendRequest(req)
	respBody := ReadResponseBody(resp)
	return ParseResponseBody(respBody, entities)
}

func CreateRequest(url string, method string) *http.Request {
	if method != "GET" && method != "POST" {
		logger.Info("Incorrect Request Method: "+method, logrus.Fields{logger.LoggerCategory: logger.LoggerCategoryHTTP})
		return nil
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logger.Info(err.Error(), logrus.Fields{logger.LoggerCategory: logger.LoggerCategoryHTTP})
	}
	return req
}

func SendRequest(req *http.Request) *http.Response {
	client := http.Client{
		Timeout: time.Second * 2,
	}
	resp, getErr := client.Do(req)
	if getErr != nil {
		logger.Fatal(getErr.Error(), logrus.Fields{logger.LoggerCategory: logger.LoggerCategoryHTTP})
	}
	return resp
}

func ReadResponseBody(resp *http.Response) []byte {
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		logger.Fatal(readErr.Error(), logrus.Fields{logger.LoggerCategory: logger.LoggerCategoryHTTP})
	}
	resp.Body.Close()
	return body
}

func ParseResponseBody(respBody []byte, entities []*Entity) []*Entity {
	jsonErr := json.Unmarshal(respBody, &entities)
	if jsonErr != nil {
		logger.Info(jsonErr.Error(), logrus.Fields{logger.LoggerCategory: logger.LoggerCategoryHTTP})
		return nil
	}
	entitiesCount := string(rune(len(entities)))
	logger.Info("Parsed "+entitiesCount+" entities.", logrus.Fields{logger.LoggerCategory: logger.LoggerCategoryHTTP})
	return entities
}
