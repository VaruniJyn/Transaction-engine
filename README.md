# Transaction Processing Engine (Go)

## Overview

This project is a backend service built in Go that simulates a simplified card transaction system.
It supports secure PIN validation, transaction processing, and transaction history tracking using in-memory storage.

---

## Features

* Card storage using in-memory database (map)
* Secure PIN hashing using SHA-256
* Transaction processing:

  * Withdraw
  * Top-up
* Validation checks:

  * Card existence
  * Card status (ACTIVE/BLOCKED)
  * PIN verification
  * Sufficient balance
* Transaction logging with unique ID and timestamp
* REST APIs for:

  * Processing transactions
  * Checking balance
  * Viewing transaction history

---

## Tech Stack

* Language: Go (Golang)
* API: `net/http` (standard library)
* Storage: In-memory (maps & slices)

---

## Setup & Run

### 1. Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/transaction-engine.git
cd transaction-engine
```

### 2. Run the server

```bash
go run .
```

Server will start at:

```
http://localhost:8080
```

---

## API Endpoints

### 1. Process Transaction

**POST** `/api/transaction`

#### Request:

```json
{
  "cardNumber": "4123456789012345",
  "pin": "1234",
  "type": "withdraw",
  "amount": 200
}
```

#### Response (Success):

```json
{
  "status": "SUCCESS",
  "respCode": "00",
  "balance": 800
}
```

#### Error Responses:

* Invalid Card → `"respCode": "05"`
* Invalid PIN → `"respCode": "06"`
* Insufficient Balance → `"respCode": "99"`

---

### 2. Get Balance

**GET** `/api/card/balance/{cardNumber}`

#### Example:

```
http://localhost:8080/api/card/balance/4123456789012345
```

#### Response:

```json
{
  "status": "SUCCESS",
  "balance": 800
}
```

---

### 3. Get Transaction History

**GET** `/api/card/transactions/{cardNumber}`

#### Example:

```
http://localhost:8080/api/card/transactions/4123456789012345
```

#### Response:

```json
[
  {
    "TransactionID": "TXN-...",
    "CardNumber": "4123456789012345",
    "Type": "withdraw",
    "Amount": 200,
    "Status": "SUCCESS",
    "Timestamp": "2026-03-25 ..."
  }
]
```

---

## API Testing (Curl)

### Transaction API

```bash
curl -X POST http://localhost:8080/api/transaction \
-H "Content-Type: application/json" \
-d '{"cardNumber":"4123456789012345","pin":"1234","type":"withdraw","amount":200}'
```

---

### Get Balance

```bash
curl http://localhost:8080/api/card/balance/4123456789012345
```

---

### Get Transaction History

```bash
curl http://localhost:8080/api/card/transactions/4123456789012345
```

---

## Security

* PINs are stored using SHA-256 hashing
* Plaintext PIN is never stored
* PIN values are never logged
* Hash comparison is used for authentication

---

## Notes

* Data is stored in-memory (resets on server restart)
* Designed as a simplified transaction processing engine

---

## Author

Varuni D
