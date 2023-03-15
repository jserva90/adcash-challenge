package models

import "time"

type LoanApplication struct {
	Amount     float64 `json:"amount"`
	Term       int     `json:"term"`
	Name       string  `json:"name"`
	PersonalID string  `json:"personal_id"`
}

type Loan struct {
	Id              int
	BorrowerName    string
	PersonalID      string
	LoanAmount      float64
	MonthlyInterest float64
	Term            int
	Status          string
	Created         time.Time
}
