package main

import (
	"encoding/base32"
	"fmt"
	"strings"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func main() {
	secret := "kirilldavidsorokin@gmail.comHENNGECHALLENGE003"
	// Encode the secret in base32
	encodedSecret := base32.StdEncoding.EncodeToString([]byte(secret))

	// Remove padding if present, and convert to upper case as required by the library
	encodedSecret = strings.TrimRight(encodedSecret, "=")
	encodedSecret = strings.ToUpper(encodedSecret)

	otp, err := totp.GenerateCodeCustom(encodedSecret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    10,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		fmt.Println("Error generating TOTP:", err)
		return
	}
	fmt.Println("Generated TOTP:", otp)
}
