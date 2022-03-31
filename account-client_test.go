package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func prepareTestAccountClient() (*AccountClient, *http.ServeMux, func()) {
	multiplexer := http.NewServeMux()
	server := httptest.NewServer(multiplexer)
	httpClient := NewHttpClient(nil)
	accountClient := NewAccountClient(httpClient)
	accountClient.HttpClient.BaseURL = server.URL + UNIT_ACCOUNTS_API_BASE
	return accountClient, multiplexer, server.Close
}

func prepareTestAccountClientWithWrongUrl() (*AccountClient, *http.ServeMux, func()) {
	multiplexer := http.NewServeMux()
	server := httptest.NewServer(multiplexer)
	httpClient := NewHttpClient(nil)
	accountClient := NewAccountClient(httpClient)
	accountClient.HttpClient.BaseURL = server.URL + ":8080"
	return accountClient, multiplexer, server.Close
}

func TestAccountClient_NewAccountClient(t *testing.T) {
	accountClient := NewAccountClient(nil)
	if accountClient != nil {
		t.Logf("SUCCESS: accountClient is found not nil %v\n", accountClient)
	} else {
		t.Errorf("FAILED: accountClient is found  nil\n")
	}
}

func TestAccountClient_FetchById(t *testing.T) {
	accountClient, multiplexer, close := prepareTestAccountClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE + "/" + SINGLE_ACCOUNT_ID
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, SINGLE_ACCOUNT_MOCK_RESPONSE)
		if err != nil {
			t.Errorf("FAILED: Error occured while writing mock response: %v\n", err)
		}
	})

	expectedAccount := populateSingleAccountDataUnitTest()

	fetchedAccount, _, res, err := accountClient.FetchById(SINGLE_ACCOUNT_ID)
	if err != nil {
		t.Errorf("FAILED: FetchById returned error: %v", err)
	}

	if !reflect.DeepEqual(fetchedAccount, expectedAccount) {
		t.Errorf("FAILED: FetchById call expected %+v, got %+v", expectedAccount, fetchedAccount)
	} else {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 200, res.StatusCode)
	}
}

func TestAccountClient_FetchById_NotFound(t *testing.T) {
	accountClient, multiplexer, close := prepareTestAccountClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE + "/" + WRONG_ACCOUNT_ID
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, err := fmt.Fprint(w, WRONG_ACCOUNT_MOCK_RESPONSE)
		if err != nil {
			t.Errorf("FAILED: Error occured while writing mock response: %v\n", err)
		}
	})

	_, _, res, err := accountClient.FetchById(WRONG_ACCOUNT_ID)
	if res.StatusCode == 404 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 404, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling FetchById: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 404, res.StatusCode)
	}
}

func TestAccountClient_ListAccount(t *testing.T) {
	accountClient, multiplexer, close := prepareTestAccountClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, MULTI_ACCOUNT_MOCK_RESPONSE)
		if err != nil {
			t.Errorf("FAILED: Error occured while writing mock response: %v\n", err)
		}
	})

	list, _, res, err := accountClient.ListAccount(nil)
	if err != nil {
		t.Errorf("FAILED: ListAccount returned error: %v", err)
	}

	if len(list) <= 0 {
		t.Errorf("FAILED: ListAccount call expected %+v, got %+v", 2, len(list))
	} else {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 200, res.StatusCode)
	}
}

func TestAccountClient_ListAccountWithWrongUrl(t *testing.T) {
	accountClient, multiplexer, close := prepareTestAccountClientWithWrongUrl()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprint(w, MULTI_ACCOUNT_MOCK_RESPONSE)
		if err != nil {
			t.Errorf("FAILED: Error occured while writing mock response: %v\n", err)
		}
	})

	_, _, _, err := accountClient.ListAccount(nil)
	if err != nil {
		t.Logf("SUCCESS: ListAccount returned error: %v", err)
	}
}

func TestAccountClient_CreateAccount(t *testing.T) {
	accountClient, multiplexer, close := prepareTestAccountClient()
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

	_, _, res, err := accountClient.CreateAccount(accountData)
	if err != nil {
		t.Errorf("FAILED: Error while calling CreateAccount: %v\n", err)
	} else if res.StatusCode == 201 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 201, res.StatusCode)
	} else {
		t.Errorf("SUCCESS: status code expected %v, got %v\n", 201, res.StatusCode)
	}
}

func TestAccountClient_CreateAccountWithWrongUrl(t *testing.T) {
	accountClient, multiplexer, close := prepareTestAccountClientWithWrongUrl()
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

	_, _, _, err := accountClient.CreateAccount(accountData)
	if err != nil {
		t.Logf("SUCCESS: CreateAccount returned error: %v", err)
	}
}

func TestAccountClient_CreateAccountWithWrongData(t *testing.T) {
	accountClient, multiplexer, close := prepareTestAccountClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, WRONG_ACCOUNT_MOCK_RESPONSE)
		if err != nil {
			t.Errorf("FAILED: Error occured while writing mock response: %v\n", err)
		}
	})

	accountData := populateWrongAccountDataUnitTest()

	_, _, res, err := accountClient.CreateAccount(accountData)
	if res.StatusCode == 400 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 400, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling CreateAccount: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 400, res.StatusCode)
	}
}

func TestAccountClient_DeleteAccount(t *testing.T) {
	accountClient, multiplexer, close := prepareTestAccountClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE + "/" + SINGLE_ACCOUNT_ID
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	res, err := accountClient.DeleteAccount(SINGLE_ACCOUNT_ID, 0)
	if err != nil {
		t.Errorf("FAILED: Error while calling DeleteAccount: %v\n", err)
	} else if res.StatusCode == 204 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 204, res.StatusCode)
	} else {
		t.Errorf("SUCCESS: status code expected %v, got %v\n", 204, res.StatusCode)
	}
}

func TestAccountClient_DeleteAccountWithWrongUrl(t *testing.T) {
	accountClient, multiplexer, close := prepareTestAccountClientWithWrongUrl()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE + "/" + SINGLE_ACCOUNT_ID
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	_, err := accountClient.DeleteAccount(SINGLE_ACCOUNT_ID, 0)
	if err != nil {
		t.Logf("SUCCESS: DeleteAccount returned error: %v", err)
	}
}

func TestAccountClient_DeleteAccount_NotFound(t *testing.T) {
	accountClient, multiplexer, close := prepareTestAccountClient()
	defer close()

	muxUrl := UNIT_ACCOUNTS_API_BASE + "/" + WRONG_ACCOUNT_ID
	multiplexer.HandleFunc(muxUrl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	res, err := accountClient.DeleteAccount(WRONG_ACCOUNT_ID, 0)
	if res.StatusCode == 404 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 404, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling DeleteAccount: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 404, res.StatusCode)
	}
}
