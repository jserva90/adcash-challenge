# Adcash Challenge

This is a simple web service that allows users to apply for a loan and view their loan history. The project was written using Golang, sqlite3 and Gin framework.

## Requirements
Python, PHP, JavaScript, Typescript, or Go programming language
Total control over frameworks, tools, and libraries
## Endpoints
The app is a headless web service that exposes the following endpoints:

Apply for a loan by providing loan amount, term, name, and personal id
List all loans by a borrower (you can skip the authorization part)
Applying for a Loan
To apply for a loan, send a POST request to the /applyforloan endpoint with a JSON payload that includes the loan amount, term, name, and personal id.

Example curl request:

curl -X POST http://localhost:8080/applyforloan \
-H "Content-Type: application/json" \
-d '{"amount": 1000, "term": 12, "name": "John Doe", "personal_id": "123456789"}'
If the loan application is approved, the endpoint will return a JSON response with loan application accepted.

If the loan application is rejected, the endpoint will return a JSON response with an error message explaining the reason for the rejection.

### Listing Loans
To list all loans for a borrower, send a GET request to the /loans/<personal_id> endpoint, where personal_id is the borrower's personal id.

Example curl request:

curl http://localhost:8080/loans/123456789
The endpoint will return a JSON response with an array of loan details.

### Lending Rules
Monthly interest rate: 5%
Reject the loan application if:
The borrower is blacklisted (store blacklisted personal ids in a config file or a database)
There have been too many applications from one personal id in the last 24 hours

### Deployment
The app can be deployed using Docker. To build the Docker image, run:
`bash docker.sh`

To delete the image, run:
`bash remove.sh`

Alternatively, you can run the app directly using the `go run .` command.

For testing, run:
`bash test.sh`

### Written in Go version 1.19
#### Authors [jserva90](https://github.com/jserva90)