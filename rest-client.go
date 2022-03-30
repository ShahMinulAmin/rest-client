package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type ResponseError struct {
	Message string `json:"error_message"`
}

type HttpClient struct {
	client  *http.Client
	BaseURL string
}

type RestClient struct {
	HttpClient *HttpClient
}

type ClientSetting struct {
	BaseURL string
	Timeout int
}

var CLIENT_SETTING_DEFAULT = &ClientSetting{
	BaseURL: "http://accountapi:8080/v1/organisation/accounts",
	Timeout: 5000,
}

func NewHttpClient(setting *ClientSetting) *HttpClient {
	if setting == nil {
		setting = CLIENT_SETTING_DEFAULT
	}
	return &HttpClient{
		client: &http.Client{
			Timeout: time.Duration(setting.Timeout) * time.Millisecond,
		},
		BaseURL: setting.BaseURL,
	}
}

func (httpClient *HttpClient) Get(url string, payload interface{}, responseData interface{}, linkData interface{}) (*http.Response, error) {
	request, err := httpClient.newHttpRequest("GET", url, payload)
	if err != nil {
		return nil, err
	}

	return httpClient.perform(context.Background(), request, responseData, linkData)
}

func (httpClient *HttpClient) Post(url string, payload interface{}, responseData interface{}, linkData interface{}) (*http.Response, error) {
	request, err := httpClient.newHttpRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	return httpClient.perform(context.Background(), request, responseData, linkData)
}

func (httpClient *HttpClient) Delete(url string) (*http.Response, error) {
	request, err := httpClient.newHttpRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return httpClient.perform(context.Background(), request, nil, nil)
}

func (httpClient *HttpClient) newHttpRequest(method, url string, bodyType interface{}) (*http.Request, error) {
	var payloadBuffer io.Reader
	if bodyType != nil {
		bodyData := ResponseBody{Data: bodyType}
		bodyJson, err := json.Marshal(bodyData)
		if err != nil {
			return nil, err
		}
		payloadBuffer = bytes.NewBuffer(bodyJson)
	}

	request, err := http.NewRequest(method, url, payloadBuffer)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	return request, nil
}

func (httpClient *HttpClient) perform(ctx context.Context, httpRequest *http.Request, responseData interface{}, linkData interface{}) (*http.Response, error) {
	httpRequest = httpRequest.WithContext(ctx)
	httpResponse, err := httpClient.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()

	responseBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	responseError := &ResponseError{}
	_ = json.Unmarshal(responseBytes, responseError)
	if responseError.Message != "" {
		return httpResponse, errors.New(responseError.Message)
	}

	if responseData != nil && linkData != nil {
		responseBody := &ResponseBody{}
		err = json.Unmarshal(responseBytes, responseBody)
		if err != nil {
			return nil, err
		}

		encodedData, err := json.Marshal(responseBody.Data)
		if err == nil {
			err = json.Unmarshal(encodedData, responseData)
			if err != nil {
				return nil, err
			}
		}

		encodedLinks, err := json.Marshal(responseBody.Links)
		if err == nil {
			err = json.Unmarshal(encodedLinks, linkData)
			if err != nil {
				return nil, err
			}
		}
	}

	return httpResponse, err
}
