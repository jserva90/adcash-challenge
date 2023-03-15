package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jserva90/adcash-challenge/helpers"
	"github.com/jserva90/adcash-challenge/models"
)

func main() {
	router := gin.Default()

	helpers.InitDatabase()

	router.POST("/applyforloan", applyForLoan)
	router.GET("/loans/:personal_id", listLoansByBorrower)

	router.Run(":8080")
}

func applyForLoan(c *gin.Context) {
	var application models.LoanApplication
	err := c.BindJSON(&application)

	if application.PersonalID == "" || application.Amount == 0 || application.Name == "" || application.Term == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "One or more of the required fields are empty."})
		return
	}

	if application.Amount < 300 || application.Amount > 10000 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Loan amount must be between 300 and 10000 euros."})
		return
	}

	if application.Term < 3 || application.Term > 60 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Loan period must be between 3 months to 5 years."})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !helpers.ValidateApplicant(application.PersonalID, c) {
		return
	}

	helpers.DB.InsertValues(application)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "The loan you have applied for has been received, for the status, check back soon."})
}

func listLoansByBorrower(c *gin.Context) {
	personalID := c.Param("personal_id")
	fmt.Println(personalID)
	getLoansWithId, err := helpers.DB.ExtractLoanData(personalID)
	helpers.ErrCheck(err)
	if len(getLoansWithId) == 0 {
		msg := fmt.Sprintf("No loans for user %s", personalID)
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.IndentedJSON(http.StatusOK, getLoansWithId)
}
