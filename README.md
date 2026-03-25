# Transaction Processing Engine (Go)

## Overview

This project is a backend service built in Go that simulates a simplified card transaction system.
It supports secure PIN validation, transaction processing, and transaction history tracking using in-memory storage.

---

## Features

* Card storage (in-memory)
* Secure PIN hashing using SHA-256
* Transaction processing:

  * Withdraw
  * Top-up
* Transaction validation:

  * Card existence
  * Card status (ACTIVE/BLOCKED)
  * PIN verification
  * Balance check
* Transaction logging with timestamp
* REST APIs for:

  * Processing transactions
  * Checking balance
  * Viewing transaction history

---

## Tech Stack

* Language: Go (Golang)
* Storage: In-memory (maps & slices)
* API: net/http (standard library)

---

## Setup & Run

```bash
go run .
```

Server runs at:

```
http://localhost:8080
```

---

## API Endpoints

### Process Transaction

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

#### Response:

```json
{
  "status": "SUCCESS",
  "respCode": "00",
  "balance": 800
}
```

---

### Get Balance

**GET** `/api/card/balance/{cardNumber}`

---

### Get Transaction History

**GET** `/api/card/transactions/{cardNumber}`

---

## Security

* PINs are stored using SHA-256 hashing
* No plaintext PIN storage
* PINs are never logged

---

## Example (PowerShell)

```powershell
Invoke-RestMethod -Method POST `
-Uri "http://localhost:8080/api/transaction" `
-ContentType "application/json" `
-Body '{"cardNumber":"4123456789012345","pin":"1234","type":"withdraw","amount":200}'
```

---

## Notes

* Data is stored in-memory (resets on server restart)
* Built as part of backend assignment
