package client

import (
	"crypto/rand"
	"fmt"
	"time"
)

const (
	INTEGRATION_ACCOUNTS_API_BASE_URL = "http://accountapi:8080/v1/organisation/accounts"
	INTEGRATION_TIME_OUT              = 5000
	UNIT_ACCOUNTS_API_BASE            = "/v1/organisation/accounts"
	SINGLE_ACCOUNT_ID                 = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	WRONG_ACCOUNT_ID                  = "adc"
	SINGLE_ACCOUNT_MOCK_RESPONSE      = `{
		"data": {
			"attributes": {
				"account_classification": "Personal",
				"account_matching_opt_out": false,
				"account_number": "10000001",
				"alternative_names": null,
				"bank_id": "400300",
				"bank_id_code": "GBDSC",
				"base_currency": "GBP",
				"bic": "NWBKGB22",
				"country": "GB",
				"iban": "GB43NWBK40030212764896",
				"joint_account": false,
				"name": [
					"Shah Minul Amin"
				],
				"secondary_identification": "X",
				"switched": false
			},
			"created_on": "2022-03-28T19:16:20.103Z",
			"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			"modified_on": "2022-03-28T19:16:20.103Z",
			"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			"type": "accounts",
			"version": 0
		},
		"links": {
			"self": "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
		}
	}`
	WRONG_ACCOUNT_MOCK_RESPONSE = `{
		"error_message": "Error occurred"
	}`
	MULTI_ACCOUNT_MOCK_RESPONSE = `{
		"data": [
			{
				"attributes": {
					"account_classification": "Personal",
					"account_matching_opt_out": false,
					"account_number": "10000001",
					"alternative_names": null,
					"bank_id": "400300",
					"bank_id_code": "GBDSC",
					"base_currency": "GBP",
					"bic": "NWBKGB22",
					"country": "GB",
					"iban": "GB28NWBK40030212764205",
					"joint_account": true,
					"name": [],
					"secondary_identification": "X",
					"switched": false
				},
				"created_on": "2022-03-28T17:39:37.171Z",
				"id": "b91afcdb-62d2-4185-b23d-71c98eaab836",
				"modified_on": "2022-03-28T17:39:37.171Z",
				"organisation_id": "b91afcdb-62d2-4185-b23d-71c98eaab817",
				"type": "accounts",
				"version": 0
			},
			{
				"attributes": {
					"account_classification": "Personal",
					"account_matching_opt_out": false,
					"account_number": "10000001",
					"alternative_names": null,
					"bank_id": "400300",
					"bank_id_code": "GBDSC",
					"base_currency": "GBP",
					"bic": "NWBKGB22",
					"country": "GB",
					"iban": "GB43NWBK40030212764896",
					"joint_account": false,
					"name": [
						"Shah Minul Amin"
					],
					"secondary_identification": "X",
					"switched": false
				},
				"created_on": "2022-03-28T18:13:20.328Z",
				"id": "515c7029-ca03-5d4d-8d33-ddf288c6cbac",
				"modified_on": "2022-03-28T18:13:20.328Z",
				"organisation_id": "b3a89ff3-2056-f93f-0050-40362eb3a7a8",
				"type": "accounts",
				"version": 0
			}
		],
		"links": {
			"first": "/v1/organisation/accounts?page%5Bnumber%5D=first&page%5Bsize%5D=2",
			"last": "/v1/organisation/accounts?page%5Bnumber%5D=last&page%5Bsize%5D=2",
			"next": "/v1/organisation/accounts?page%5Bnumber%5D=1&page%5Bsize%5D=2",
			"self": "/v1/organisation/accounts?page%5Bnumber%5D=0&page%5Bsize%5D=2"
		}
	}`
)

func uuid() string {
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

func populateAccountAttributes() *AccountAttributes {
	classification := "Personal"
	matchingOptOut := false
	country := "GB"
	joint := false
	switched := false
	names := []string{"Shah Minul Amin"}
	attributes := &AccountAttributes{
		AccountClassification:   &classification,
		AccountMatchingOptOut:   &matchingOptOut,
		AccountNumber:           "10000001",
		BankID:                  "400300",
		BankIDCode:              "GBDSC",
		BaseCurrency:            "GBP",
		Bic:                     "NWBKGB22",
		Country:                 &country,
		Iban:                    "GB43NWBK40030212764896",
		JointAccount:            &joint,
		SecondaryIdentification: "X",
		Switched:                &switched,
		Name:                    names,
	}
	return attributes
}

func populateSingleAccountDataUnitTest() *AccountData {
	attributes := populateAccountAttributes()
	orgId := "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	var version int64 = 0
	timestamp := time.Date(2022, time.March, 28, 19, 16, 20, 103000000, time.UTC)
	accountData := &AccountData{
		ID:             SINGLE_ACCOUNT_ID,
		OrganisationID: orgId,
		Type:           "accounts",
		Version:        &version,
		CreatedOn:      &timestamp,
		ModifiedOn:     &timestamp,
		Attributes:     attributes,
	}
	return accountData
}

func populateWrongAccountDataUnitTest() *AccountData {
	attributes := populateAccountAttributes()
	orgId := "ebf"
	var version int64 = 0
	timestamp := time.Date(2022, time.March, 28, 19, 16, 20, 103000000, time.UTC)
	accountData := &AccountData{
		ID:             WRONG_ACCOUNT_ID,
		OrganisationID: orgId,
		Type:           "accounts",
		Version:        &version,
		CreatedOn:      &timestamp,
		ModifiedOn:     &timestamp,
		Attributes:     attributes,
	}
	return accountData
}

func populateAccountDataIntegration() *AccountData {
	attributes := populateAccountAttributes()
	id := uuid()
	orgId := uuid()
	var version int64 = 0
	currentTime := time.Now()
	accountData := &AccountData{
		ID:             id,
		OrganisationID: orgId,
		Type:           "accounts",
		Version:        &version,
		CreatedOn:      &currentTime,
		ModifiedOn:     &currentTime,
		Attributes:     attributes,
	}
	return accountData
}
