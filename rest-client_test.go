package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func prepareTestRestClient() (*HttpClient, *http.ServeMux, func()) {
	multiplexer := http.NewServeMux()
	server := httptest.NewServer(multiplexer)
	httpClient := NewHttpClient(nil)
	httpClient.BaseURL = server.URL + UNIT_ACCOUNTS_API_BASE
	return httpClient, multiplexer, server.Close
}

func TestRestClient_Get(t *testing.T) {
	httpClient, multiplexer, close := prepareTestRestClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE + "/" + SINGLE_ACCOUNT_ID
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, SINGLE_ACCOUNT_MOCK_RESPONSE)
		if err != nil {
			t.Errorf("FAILED: Error occured while writing mock response: %v\n", err)
		}
	})

	expectedResponse := populateSingleAccountDataUnitTest()
	accountResponse := new(AccountData)
	links := new(Links)
	url := fetchAccountApiUrl(httpClient.BaseURL, SINGLE_ACCOUNT_ID)

	res, err := httpClient.Get(url, nil, accountResponse, links)
	if err != nil {
		t.Errorf("FAILED: httpClient Get returned error: %v", err)
	}

	if !reflect.DeepEqual(accountResponse, expectedResponse) {
		t.Errorf("FAILED: httpClient Get call expected %+v, got %+v", expectedResponse, accountResponse)
	} else {
		t.Logf("SUCCESS: http status code expected %v, got %v\n", 200, res.StatusCode)
	}
}

func TestRestClient_GetWithWrongUrl(t *testing.T) {
	httpClient, multiplexer, close := prepareTestRestClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE + "/" + SINGLE_ACCOUNT_ID
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, SINGLE_ACCOUNT_MOCK_RESPONSE)
		if err != nil {
			t.Errorf("FAILED: Error occured while writing mock response: %v\n", err)
		}
	})

	accountResponse := new(AccountData)
	links := new(Links)
	url := ":8080"
	_, err := httpClient.Get(url, nil, accountResponse, links)
	if err != nil {
		t.Logf("SUCCESS: httpClient Get returned error: %v", err)
	}
}

func TestRestClient_Post(t *testing.T) {
	httpClient, multiplexer, close := prepareTestRestClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, err := fmt.Fprint(w, SINGLE_ACCOUNT_MOCK_RESPONSE)
		if err != nil {
			t.Errorf("FAILED: Error occured while writing mock response: %v\n", err)
		}
	})

	accountData := populateSingleAccountDataUnitTest()
	accountResponse := new(AccountData)
	links := new(Links)
	res, err := httpClient.Post(httpClient.BaseURL, accountData, accountResponse, links)

	if err != nil {
		t.Errorf("FAILED: Error while calling Post: %v\n", err)
	} else if res.StatusCode == 201 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 201, res.StatusCode)
	} else {
		t.Errorf("SUCCESS: status code expected %v, got %v\n", 201, res.StatusCode)
	}
}

func TestRestClient_PostWithWrongUrl(t *testing.T) {
	httpClient, multiplexer, close := prepareTestRestClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE + "/" + SINGLE_ACCOUNT_ID
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, SINGLE_ACCOUNT_MOCK_RESPONSE)
		if err != nil {
			t.Errorf("FAILED: Error occured while writing mock response: %v\n", err)
		}
	})

	accountResponse := new(AccountData)
	links := new(Links)
	url := ":8080"
	_, err := httpClient.Post(url, nil, accountResponse, links)
	if err != nil {
		t.Logf("SUCCESS: httpClient Post returned error: %v", err)
	}
}

func TestRestClient_Delete(t *testing.T) {
	httpClient, multiplexer, close := prepareTestRestClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE + "/" + SINGLE_ACCOUNT_ID
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	url := deleteAccountApiUrl(httpClient.BaseURL, SINGLE_ACCOUNT_ID, 0)
	res, err := httpClient.Delete(url)

	if err != nil {
		t.Errorf("FAILED: Error while calling DeleteAccount: %v\n", err)
	} else if res.StatusCode == 204 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 204, res.StatusCode)
	} else {
		t.Errorf("SUCCESS: status code expected %v, got %v\n", 204, res.StatusCode)
	}
}

func TestRestClient_DeleteWithWrongUrl(t *testing.T) {
	httpClient, multiplexer, close := prepareTestRestClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE + "/" + SINGLE_ACCOUNT_ID
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	url := ":8080"
	_, err := httpClient.Delete(url)

	if err != nil {
		t.Logf("SUCCESS: httpClient Delete returned error: %v", err)
	}
}
