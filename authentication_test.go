package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	main "WeVentureTask"
)

const loginUrl = "http://localhost:8080" + main.BaseURL + main.LoginPath

func TestSuccessfulLogin(t *testing.T) {
	body, _ := json.Marshal(map[string]string{
		"username": "mostafa",
		"password": "testtest",
	})

	resp, err := http.Post(loginUrl,"application/json", bytes.NewReader(body))

	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Status should be %d but it is %d", http.StatusOK, resp.StatusCode)
	}

	var respBody map[string]string
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Fatal(err)
	}

	if val, ok := respBody["token"]; !ok || val == "" {
		t.Fatal("no token returned")
	}
}

func TestBadRequestLogin(t *testing.T) {
	body, _ := json.Marshal(map[string]string{
		"username": "mostafa",
	})

	resp, err := http.Post(loginUrl, "application/json", bytes.NewReader(body))

	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Status should be %d but it is %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestUnauthorizedLogin(t *testing.T) {
	body := []map[string]string{
		{
			"username": "mostaf",
			"password": "testtest",
		},
		{
			"username": "mostafa",
			"password": "testtes",
		},
	}

	for _, tc := range body {
		reqBody, _ := json.Marshal(tc)

		resp, err := http.Post(loginUrl, "application/json", bytes.NewReader(reqBody))

		//Handle Error
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}

		if resp.StatusCode != http.StatusUnauthorized {
			t.Fatalf("Status should be %d but it is %d", http.StatusUnauthorized, resp.StatusCode)
		}

		resp.Body.Close()
	}
}
