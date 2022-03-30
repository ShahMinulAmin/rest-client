package client

import (
	"net/http"
	"testing"
)

func prepareClient() *AccountClient {
	setting := &ClientSetting{
		BaseURL: INTEGRATION_ACCOUNTS_API_BASE_URL,
		Timeout: INTEGRATION_TIME_OUT,
	}
	httpClient := NewHttpClient(setting)
	accountClient := NewAccountClient(httpClient)
	return accountClient
}

func createAccount() (*AccountData, *Links, *http.Response, error) {
	accountClient := prepareClient()
	accountData := populateAccountDataIntegration()
	return accountClient.CreateAccount(accountData)

}

func createAccounts(count int) {
	for i := 0; i < count; i++ {
		createAccount()
	}
}

func TestCreateAccount_CorrectAccountData(t *testing.T) {
	_, _, res, err := createAccount()

	if err != nil {
		t.Errorf("FAILED: Error while calling CreateAccount: %v\n", err)
	} else if res.StatusCode == 201 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 201, res.StatusCode)
	} else {
		t.Errorf("SUCCESS: status code expected %v, got %v\n", 201, res.StatusCode)
	}
}

func TestCreateAccount_IncorrectID(t *testing.T) {
	accountClient := prepareClient()
	accountData := populateAccountDataIntegration()
	accountData.ID = "abc"
	_, _, res, err := accountClient.CreateAccount(accountData)

	if res.StatusCode == 400 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 400, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling CreateAccount: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 400, res.StatusCode)
	}
}

func TestCreateAccount_IncorrectOrganisationID(t *testing.T) {
	accountClient := prepareClient()
	accountData := populateAccountDataIntegration()
	accountData.OrganisationID = "abc"
	_, _, res, err := accountClient.CreateAccount(accountData)

	if res.StatusCode == 400 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 400, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling CreateAccount: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 400, res.StatusCode)
	}
}

func TestCreateAccount_NoAttributes(t *testing.T) {
	accountClient := prepareClient()
	accountData := populateAccountDataIntegration()
	accountData.Attributes = nil
	_, _, res, err := accountClient.CreateAccount(accountData)

	if res.StatusCode == 400 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 400, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling CreateAccount: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 400, res.StatusCode)
	}
}

func TestCreateAccount_NoName(t *testing.T) {
	accountClient := prepareClient()
	accountData := populateAccountDataIntegration()
	accountData.Attributes.Name = nil
	_, _, res, err := accountClient.CreateAccount(accountData)

	if res.StatusCode == 400 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 400, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling CreateAccount: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 400, res.StatusCode)
	}
}

func TestCreateAccount_IncorrectAccountClassification(t *testing.T) {
	accountClient := prepareClient()
	accountData := populateAccountDataIntegration()
	classification := "A"
	accountData.Attributes.AccountClassification = &classification
	_, _, res, err := accountClient.CreateAccount(accountData)

	if res.StatusCode == 400 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 400, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling CreateAccount: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 400, res.StatusCode)
	}
}

func TestCreateAccount_IncorrectCountry(t *testing.T) {
	accountClient := prepareClient()
	accountData := populateAccountDataIntegration()
	country := "A"
	accountData.Attributes.Country = &country
	_, _, res, err := accountClient.CreateAccount(accountData)

	if res.StatusCode == 400 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 400, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling CreateAccount: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 400, res.StatusCode)
	}
}

func TestCreateAccount_IncorrectBaseCurrency(t *testing.T) {
	accountClient := prepareClient()
	accountData := populateAccountDataIntegration()
	accountData.Attributes.BaseCurrency = "A"
	_, _, res, err := accountClient.CreateAccount(accountData)

	if res.StatusCode == 400 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 400, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling CreateAccount: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 400, res.StatusCode)
	}
}

func TestDeleteAccount(t *testing.T) {
	accountClient := prepareClient()
	accountData := populateAccountDataIntegration()
	accountClient.CreateAccount(accountData)
	res, err := accountClient.DeleteAccount(accountData.ID, int(*accountData.Version))

	if err != nil {
		t.Errorf("FAILED: Error while calling DeleteAccount: %v\n", err)
	} else if res.StatusCode == 204 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 204, res.StatusCode)
	} else {
		t.Errorf("SUCCESS: status code expected %v, got %v\n", 204, res.StatusCode)
	}
}

func TestDeleteAccount_NotFound(t *testing.T) {
	accountClient := prepareClient()
	id := uuid()
	res, err := accountClient.DeleteAccount(id, 0)

	if res.StatusCode == 404 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 404, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling DeleteAccount: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 404, res.StatusCode)
	}
}

func TestFetchById(t *testing.T) {
	accountClient := prepareClient()
	accountData := populateAccountDataIntegration()
	accountClient.CreateAccount(accountData)
	_, _, res, err := accountClient.FetchById(accountData.ID)

	if err != nil {
		t.Errorf("FAILED: Error while calling FetchById: %v\n", err)
	} else {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 200, res.StatusCode)
	}
}

func TestFetchById_NotFound(t *testing.T) {
	accountClient := prepareClient()
	id := uuid()
	_, _, res, err := accountClient.FetchById(id)

	if res.StatusCode == 404 {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 404, res.StatusCode)
	} else if err != nil {
		t.Errorf("FAILED: Error while calling FetchById: %v\n", err)
	} else {
		t.Errorf("FAILED: status code expected %v, got %v\n", 404, res.StatusCode)
	}
}

func TestListAccount_NoParams(t *testing.T) {
	createAccounts(2)

	accountClient := prepareClient()
	_, _, res, err := accountClient.ListAccount(nil)

	if err != nil {
		t.Errorf("FAILED: Error while calling ListAccount: %v\n", err)
	} else {
		t.Logf("SUCCESS: status code expected %v, got %v\n", 200, res.StatusCode)
	}
}

func TestListAccount_WithParams(t *testing.T) {
	createAccounts(2)

	accountClient := prepareClient()
	params := &AccountParams{
		Number: "0",
		Size:   1,
	}
	accounts, _, _, err := accountClient.ListAccount(params)

	if err != nil {
		t.Errorf("FAILED: Error while calling ListAccount: %v\n", err)
	} else {
		length := len(accounts)
		if length == 1 {
			t.Logf("SUCCESS: number of accounts expected %d, got %d\n", 1, length)
		} else {
			t.Errorf("FAILED: number of accounts expected %d, got %d\n", 1, length)
		}
	}
}
