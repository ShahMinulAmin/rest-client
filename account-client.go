package client

import (
	"fmt"
	"log"
	"net/http"
)

type AccountParams struct {
	Number string
	Size   int
}

type AccountClient RestClient

// Creates a new account client using http client.
func NewAccountClient(httpClient *HttpClient) *AccountClient {
	if httpClient == nil {
		httpClient = NewHttpClient(nil)
	}
	return &AccountClient{
		HttpClient: httpClient,
	}
}

// Gets a single account using the account ID.
// Returns account data, links and http response.
func (accountClient *AccountClient) FetchById(id string) (*AccountData, *Links, *http.Response, error) {
	accountResponse := new(AccountData)
	links := new(Links)
	url := fetchAccountApiUrl(accountClient.HttpClient.BaseURL, id)

	httpResponse, err := accountClient.HttpClient.Get(url, nil, accountResponse, links)
	if err != nil {
		log.Printf("Error occurred while fetching account by id: %v\n", err)
		return nil, nil, httpResponse, err
	}

	return accountResponse, links, httpResponse, nil
}

// List accounts with optional page parameters.
// Returns list of accounts' data, links and http response.
func (accountClient *AccountClient) ListAccount(params *AccountParams) ([]*AccountData, *Links, *http.Response, error) {
	accounts := new([]*AccountData)
	links := new(Links)
	url := listAccountApiUrl(accountClient.HttpClient.BaseURL, params)

	httpResponse, err := accountClient.HttpClient.Get(url, nil, accounts, links)
	if err != nil {
		log.Printf("Error occurred while fetching account list: %v\n", err)
		return nil, nil, httpResponse, err
	}

	return *accounts, links, httpResponse, nil
}

// Creates a bank account with provided account data payload.
// Returns account data, links and http response.
func (accountClient *AccountClient) CreateAccount(payload *AccountData) (*AccountData, *Links, *http.Response, error) {
	accountResponse := new(AccountData)
	links := new(Links)
	httpResponse, err := accountClient.HttpClient.Post(accountClient.HttpClient.BaseURL, payload, accountResponse, links)
	if err != nil {
		log.Printf("Error occurred while creating account: %v\n", err)
		return nil, nil, httpResponse, err
	}

	return accountResponse, links, httpResponse, nil
}

// Deletes a account using the account ID and version number.
// Returns http response.
func (accountClient *AccountClient) DeleteAccount(id string, version int) (*http.Response, error) {
	url := deleteAccountApiUrl(accountClient.HttpClient.BaseURL, id, version)

	httpResponse, err := accountClient.HttpClient.Delete(url)
	if err != nil {
		log.Printf("Error occurred while deleting account: %v\n", err)
		return httpResponse, err
	}

	return httpResponse, nil
}

// Populates fetch account API URL from base URL and account ID
func fetchAccountApiUrl(baseURL string, id string) string {
	url := fmt.Sprintf("%s/%s", baseURL, id)
	return url
}

// Populates list account API URL from base URL and page parameters
func listAccountApiUrl(baseURL string, params *AccountParams) string {
	url := baseURL
	if params != nil {
		url = fmt.Sprintf("%s?page[number]=%s&page[size]=%d", baseURL, params.Number, params.Size)
	}
	return url
}

// Populates delete account API URL from base URL, account ID and version
func deleteAccountApiUrl(baseURL string, id string, version int) string {
	url := fmt.Sprintf("%s/%s?version=%d", baseURL, id, version)
	return url
}
