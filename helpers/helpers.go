package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func ErrCheck(inputErr error) {
	if inputErr != nil && inputErr != sql.ErrNoRows {
		// Open the log file for appending, create if it doesn't exist
		f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()

		// Add a timestamp to the error message
		errorMsg := fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), inputErr.Error())

		// Write the error message to the log file
		if _, err = f.WriteString(errorMsg + "\n"); err != nil {
			log.Println(err)
		}

		log.Println("See error:", inputErr)
	}
}
