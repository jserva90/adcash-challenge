#!/bin/bash

# Test loan application API
printf "\n\nTesting loan application API...\n\n"
read -n 1 -s -r -p "Press any key to continue..."

# Test successful request
printf "\n\nTesting successful request...\n\n"
curl -X POST http://localhost:8080/applyforloan -H "Content-Type: application/json" -d '{"amount": 1000, "term": 12, "name": "Jim Jones", "personal_id": "123321"}'
read -n 1 -s -r -p "Press any key to continue..."


# Test failed request (too many applications in 24h)
printf "\n\nTesting failed request...(maximum amount of applications exceeded)\n\n"
curl -X POST http://localhost:8080/applyforloan -H "Content-Type: application/json" -d '{"amount": 1000, "term": 12, "name": "Jim Jones", "personal_id": "123321"}'
read -n 1 -s -r -p "\Press any key to continue..."


# Test failed request (user is blacklisted)
printf "\n\nTesting failed request...(user is in the blacklist)\n\n"
curl -X POST http://localhost:8080/applyforloan -H "Content-Type: application/json" -d '{"amount": 1000, "term": 12, "name": "Evan Fraudson", "personal_id": "999999999"}'
read -n 1 -s -r -p "Press any key to continue..."

# Test loan listing API
printf "\n\nTesting loan listing API...\n\n"

# Test successful request
printf "\n\nTesting successful request...\n\n"
curl http://localhost:8080/loans/123456789
read -n 1 -s -r -p "Press any key to continue..."

# Test failed request
printf "\n\nTesting failed request...(user has no loans)\n\n"
curl http://localhost:8080/loans/00
read -n 1 -s -r -p "Press any key to exit..."