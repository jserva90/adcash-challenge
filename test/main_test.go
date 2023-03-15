package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jserva90/adcash-challenge/models"
)

func TestLoanApplication_MarshalJSON(t *testing.T) {
	loanApp := models.LoanApplication{
		Amount:     10000.0,
		Term:       12,
		Name:       "John Doe",
		PersonalID: "123-45-6789",
	}

	expectedJSON := `{"amount":10000,"term":12,"name":"John Doe","personal_id":"123-45-6789"}`

	jsonData, err := json.Marshal(loanApp)
	if err != nil {
		t.Fatalf("failed to marshal loan application: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("unexpected JSON output, expected %s but got %s", expectedJSON, string(jsonData))
	}
}

func TestLoan_MarshalJSON(t *testing.T) {
	testTime := getTestTime()
	loan := models.Loan{
		Id:              1,
		BorrowerName:    "John Doe",
		PersonalID:      "123-45-6789",
		LoanAmount:      10000.0,
		MonthlyInterest: 1.0,
		Term:            12,
		Status:          "approved",
		Created:         testTime,
	}

	expectedJSON := `{"Id":1,"BorrowerName":"John Doe","PersonalID":"123-45-6789","LoanAmount":10000,"MonthlyInterest":1,"Term":12,"Status":"approved","Created":"` + testTime.Format(time.RFC3339) + `"}`

	jsonData, err := json.Marshal(loan)
	if err != nil {
		t.Fatalf("failed to marshal loan: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("unexpected JSON output, expected %s but got %s", expectedJSON, string(jsonData))
	}
}

func TestLoan_MarshalJSON_ZeroTime(t *testing.T) {
	testTime := time.Time{}
	loan := models.Loan{
		Id:              1,
		BorrowerName:    "John Doe",
		PersonalID:      "123-45-6789",
		LoanAmount:      10000.0,
		MonthlyInterest: 1.0,
		Term:            12,
		Status:          "approved",
		Created:         testTime,
	}

	expectedJSON := `{"Id":1,"BorrowerName":"John Doe","PersonalID":"123-45-6789","LoanAmount":10000,"MonthlyInterest":1,"Term":12,"Status":"approved","Created":"` + testTime.Format(time.RFC3339) + `"}`

	jsonData, err := json.Marshal(loan)
	if err != nil {
		t.Fatalf("failed to marshal loan: %v", err)
	}

	if string(jsonData) != expectedJSON {
		t.Errorf("unexpected JSON output, expected %s but got %s", expectedJSON, string(jsonData))
	}
}

func getTestTime() time.Time {
	testTime, _ := time.Parse(time.RFC3339, "2023-03-15T12:00:00-07:00")
	return testTime
}
