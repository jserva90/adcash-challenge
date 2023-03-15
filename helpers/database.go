package helpers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jserva90/adcash-challenge/models"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Conn *sql.DB
}

var DB Database

func InitDatabase() {
	var err error
	dbPath := "./data/loans.db"
	exists := true

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		exists = false
	}

	DB.Conn, err = sql.Open("sqlite3", dbPath)
	ErrCheck(err)

	if !exists {
		pre := "./data/"
		paths := []string{pre + "tables.sql", pre + "mockdata.sql"}

		for _, path := range paths {
			readFile, err := os.ReadFile(path)
			ErrCheck(err)
			fileStr := string(readFile)
			_, err = DB.Conn.Exec(fileStr)
			ErrCheck(err)
		}
	}
}

func (db *Database) InsertValues(application models.LoanApplication) {
	db.Conn.Exec("INSERT INTO loan_applications (personal_id) VALUES(?)", application.PersonalID)
	db.Conn.Exec("INSERT INTO loans (borrower_name, personal_id, loan_amount, term) VALUES(?,?,?,?)", application.Name, application.PersonalID, application.Amount, application.Term)
}

func (db *Database) ExtractLoanData(personalId string) ([]models.Loan, error) {
	var loans []models.Loan
	stmt, err := db.Conn.Prepare("SELECT * FROM loans WHERE personal_id = ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(personalId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var loan models.Loan
		rows.Scan(&loan.Id, &loan.BorrowerName, &loan.PersonalID, &loan.LoanAmount, &loan.MonthlyInterest, &loan.Term, &loan.Status, &loan.Created)
		loans = append(loans, loan)
	}
	return loans, nil
}

func (db *Database) CheckForBlackList(personalId string) (bool, error) {
	stmt, err := db.Conn.Prepare("SELECT COUNT(*) FROM blacklist WHERE personal_id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(personalId).Scan(&count)
	if err != nil {
		return false, err
	}

	fmt.Println(count > 0)
	return count > 0, nil
}

func (db *Database) checkIfApplicantExceededsMaxApplicationsPerDay(personalId string) (bool, error) {
	now := time.Now()
	stmt, err := db.Conn.Prepare("SELECT created_at FROM loan_applications WHERE personal_id = ?")
	if err != nil {
		return false, err
	}
	rows, err := stmt.Query(personalId)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		var applicationDate time.Time
		rows.Scan(&applicationDate)
		dur := now.Sub(applicationDate)
		if dur < 24*time.Hour {
			return true, nil
		}
	}

	return false, nil
}

func ValidateApplicant(personalId string, c *gin.Context) bool {
	isBlackListed, err := DB.CheckForBlackList(personalId)
	ErrCheck(err)
	if isBlackListed {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Applicant is blacklisted."})
		return false
	}

	exceeds, err := DB.checkIfApplicantExceededsMaxApplicationsPerDay(personalId)
	ErrCheck(err)
	if exceeds {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Already applied for loan today."})
		return false
	}

	return true
}
