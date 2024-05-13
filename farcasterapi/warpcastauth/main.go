package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
	AccountPrivKey string
	FID            uint64
)

func main() {
	AccountPrivKey = os.Getenv("PRIVKEY")
	FIDstr := os.Getenv("FID")
	if AccountPrivKey == "" || FIDstr == "" {
		log.Fatal("PRIVKEY and FID environment variables are required")
	}
	var err error
	FID, err = strconv.ParseUint(os.Getenv("FID"), 10, 64)
	if err != nil {
		log.Fatalf("Error parsing FID: %v", err)
	}
	http.HandleFunc("/", handleRequest)
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Generate Ed25519 key pair
	signerPubKey, signerPrivKey, err := GenerateKeyPair()
	if err != nil {
		http.Error(w, "Error generating key pair", http.StatusInternalServerError)
		return
	}

	privKey, err := crypto.HexToECDSA(AccountPrivKey)
	if err != nil {
		http.Error(w, "Error converting private key", http.StatusInternalServerError)
		return
	}

	// Make API request
	deeplinkUrl, err := CreateSignedKeyRequest(privKey, signerPubKey, FID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error making signed key request: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Deep link URL: %s\n", deeplinkUrl)

	// Generate the QR code
	qrCode, err := GenerateQRCode(deeplinkUrl)
	if err != nil {
		http.Error(w, "Error generating QR code", http.StatusInternalServerError)
		return
	}

	// Display the information
	fmt.Fprintf(w, "<h1>Warpcast API Integration</h1>")
	fmt.Fprintf(w, "<p><strong>URL with Token:</strong> <a href='%s'>%s</a></p>", deeplinkUrl, deeplinkUrl)
	fmt.Fprintf(w, "<p><strong>Public Key:</strong> %x</p>", signerPrivKey)
	fmt.Fprintf(w, "<p><strong>Private Key:</strong> %x</p>", signerPubKey)
	fmt.Fprintf(w, "<p><strong>QR Code:</strong><br><img src='data:image/png;base64,%s'/></p>", base64.StdEncoding.EncodeToString(qrCode))
}
