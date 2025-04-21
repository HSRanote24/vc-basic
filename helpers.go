// helpers.go
// Contains helper functions for onboarding issuer, issuing credentials, and verifying credentials.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// sendJSONRequest is a helper function to marshal a payload, send a POST request, and return the response or error.
func sendJSONRequest(url string, payload map[string]interface{}) (string, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return string(body), fmt.Errorf("request failed: status %d, body: %s", resp.StatusCode, string(body))
	}

	return string(body), nil
}

// OnboardIssuer sends a request to onboard a new issuer and returns the response body or error.
func OnboardIssuer() (string, error) {
	url := "http://localhost:7002/onboard/issuer"
	payload := map[string]interface{}{
		"key": map[string]interface{}{
			"backend": "jwk",
			"keyType": "Ed25519",
		},
		"did": map[string]interface{}{
			"method": "jwk",
		},
	}
	return sendJSONRequest(url, payload)
}

// IssueCredential sends a request to issue a new credential and returns the response body or error.
func IssueCredential() (string, error) {
	url := "http://localhost:7002/openid4vc/jwt/issue"
	payload := map[string]interface{}{
		"issuerKey": map[string]interface{}{
			"type": "jwk",
			"jwk": map[string]interface{}{
				"kty": "OKP",
				"d":   "mDhpwaH6JYSrD2Bq7Cs-pzmsjlLj4EOhxyI-9DM1mFI",
				"crv": "Ed25519",
				"kid": "Vzx7l5fh56F3Pf9aR3DECU5BwfrY6ZJe05aiWYWzan8",
				"x":   "T3T4-u1Xz3vAV2JwPNxWfs4pik_JLiArz_WTCvrCFUM",
			},
		},
		"credentialConfigurationId": "UniversityDegree_jwt_vc_json",
		"credentialData": map[string]interface{}{
			"@context": []interface{}{
				"https://www.w3.org/2018/credentials/v1",
				"https://www.w3.org/2018/credentials/examples/v1",
			},
			"id": "http://example.gov/credentials/3732",
			"type": []interface{}{
				"VerifiableCredential",
				"UniversityDegreeCredential",
			},
			"issuer": map[string]interface{}{
				"id": "did:web:vc.transmute.world",
			},
			"issuanceDate": "2020-03-10T04:24:12.164Z",
			"credentialSubject": map[string]interface{}{
				"id": "did:example:ebfeb1f712ebc6f1c276e12ec21",
				"degree": map[string]interface{}{
					"type": "BachelorDegree",
					"name": "Bachelor of Science and Arts",
				},
			},
		},
		"mapping": map[string]interface{}{
			"id": "<uuid>",
			"issuer": map[string]interface{}{
				"id": "<issuerDid>",
			},
			"credentialSubject": map[string]interface{}{
				"id": "<subjectDid>",
			},
			"issuanceDate":   "<timestamp>",
			"expirationDate": "<timestamp-in:365d>",
		},
		"authenticationMethod": "PRE_AUTHORIZED",
		"issuerDid":            "did:key:z6MkjoRhq1jSNJdLiruSXrFFxagqrztZaXHqHGUTKJbcNywp",
		"standardVersion":      "DRAFT13",
	}
	return sendJSONRequest(url, payload)
}

// VerifyCredential sends a request to verify a credential and returns the response body or error.
func VerifyCredential() (string, error) {
	url := "http://localhost:7003/openid4vc/verify"
	payload := map[string]interface{}{
		"request_credentials": []interface{}{
			map[string]interface{}{
				"format": "jwt_vc_json",
				"type":   "OpenBadgeCredential",
			},
		},
	}
	return sendJSONRequest(url, payload)
}
