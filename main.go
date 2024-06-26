package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	jwtsigner "github.com/pantheon-systems/pantheon-jwt-signer"
)

func main() {
	ctx := context.Background()

	// TODO: Use dynamic name
	var jsondata = []byte(`{"name": "4832107511", "value": "test"}`)

	req, err := http.NewRequest(
		"POST",
		"https://customer-secrets-jwt.svc.pantheon.io/users/577adc84-4dc6-4149-b56e-ae3727395f44/secrets",
		bytes.NewBuffer(jsondata),
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Pantheon-User", "577adc84-4dc6-4149-b56e-ae3727395f44")

	// TODO: Read email/SA from config
	jwts, err := jwtsigner.NewJwtSigner(ctx, "sachin.prasad@pantheon.io")
	if err != nil {
		log.Fatal(err)
	}

	// Do i need to pass customer-secrets as audience? I believe not and passing "" will set pantheon as default audience
	err = jwts.SignRequest(ctx, req, "customer-secrets")
	if err != nil {
		log.Fatal(err)
	}

	// Save data
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Printf("Response: %v\n", resp.Body)

	read()
}

func read() {
	ctx := context.Background()
	req, err := http.NewRequest(
		"GET",
		"https://customer-secrets-jwt.svc.pantheon.io/users/577adc84-4dc6-4149-b56e-ae3727395f44/secrets/4832107511",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Pantheon-User", "577adc84-4dc6-4149-b56e-ae3727395f44")
	// TODO: Read email/SA from config
	jwts, err := jwtsigner.NewJwtSigner(ctx, "sachin.prasad@pantheon.io")
	if err != nil {
		log.Fatal(err)
	}

	// Do i need to pass customer-secrets as audience? I believe not and passing "" will set pantheon as default audience
	err = jwts.SignRequest(ctx, req, "customer-secrets")
	if err != nil {
		log.Fatal(err)
	}

	// Get data
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	// Read response body
	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)
	fmt.Printf("Response: %v\n", string(buf.Bytes()))
}
