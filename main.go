// main.go
// Entry point for the VC Basic project. This file contains functions to onboard an issuer,
// issue a credential, and verify a credential using HTTP requests to local endpoints.

package main

import (
	"fmt"
	"log"
)

// main orchestrates the onboarding, issuance, and verification steps.
func main() {
	fmt.Println("[INFO] Onboarding issuer via /onboard/issuer ...")
	result, err := OnboardIssuer()
	if err != nil {
		log.Fatalf("[ERROR] Onboarding failed: %v\nResponse: %s", err, result)
	}
	fmt.Println("[SUCCESS] Onboarded issuer. Response:")
	fmt.Println(result)

	fmt.Println("[INFO] Issuing credential via /openid4vc/jwt/issue ...")
	credResult, credErr := IssueCredential()
	if credErr != nil {
		log.Fatalf("[ERROR] Credential issuance failed: %v\nResponse: %s", credErr, credResult)
	}
	fmt.Println("[SUCCESS] Credential issued. Response:")
	fmt.Println(credResult)

	fmt.Println("[INFO] Verifying credential via /openid4vc/verify ...")
	verifyResult, verifyErr := VerifyCredential()
	if verifyErr != nil {
		log.Fatalf("[ERROR] Credential verification failed: %v\nResponse: %s", verifyErr, verifyResult)
	}
	fmt.Println("[SUCCESS] Credential verified. Response:")
	fmt.Println(verifyResult)
}
