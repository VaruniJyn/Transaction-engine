package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func generateID() string {
	return fmt.Sprintf("TXN-%d", time.Now().UnixNano())
}

func hashPIN(pin string) string {
	hash := sha256.Sum256([]byte(pin))
	return hex.EncodeToString(hash[:])
}
