package main_test

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	main "WeVentureTask"
	"github.com/dgrijalva/jwt-go"
)

const url = "http://localhost:8080" + main.BaseURL

func TestEndPointAWithRoleUserA(t *testing.T) {
	loginToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"role": "UserA",
		"exp":  time.Now().UTC().Add(time.Hour * 24).Unix(),
		"nano": time.Now().Nanosecond(),
	})

	token, err := loginToken.SignedString([]byte("SecretMostafa"))
	r, err := http.NewRequest("GET", url + "/hello/Mostafa",
		bytes.NewReader([]byte("")))
	r.Header["Authorization"] = []string{"Bearer " + token}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Status should be %d but it is %d", http.StatusOK, resp.StatusCode)
	}
}

func TestEndPointAWithRoleUserB(t *testing.T) {
	loginToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"role": "UserB",
		"exp":  time.Now().UTC().Add(time.Hour * 24).Unix(),
		"nano": time.Now().Nanosecond(),
	})

	token, err := loginToken.SignedString([]byte("SecretMostafa"))
	r, err := http.NewRequest("GET", url + "/hello/Mostafa",
		bytes.NewReader([]byte("")))
	r.Header["Authorization"] = []string{"Bearer " + token}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Status should be %d but it is %d", http.StatusOK, resp.StatusCode)
	}
}

func TestEndPointAWithNonExistRole(t *testing.T) {
	loginToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"role": "NoRole",
		"exp":  time.Now().UTC().Add(time.Hour * 24).Unix(),
		"nano": time.Now().Nanosecond(),
	})

	token, err := loginToken.SignedString([]byte("SecretMostafa"))
	r, err := http.NewRequest("GET", url + "/hello/Mostafa",
		bytes.NewReader([]byte("")))
	r.Header["Authorization"] = []string{"Bearer " + token}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("Status should be %d but it is %d", http.StatusForbidden, resp.StatusCode)
	}
}

func TestEndPointBWithRoleUserA(t *testing.T) {
	loginToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"role": "UserA",
		"exp":  time.Now().UTC().Add(time.Hour * 24).Unix(),
		"nano": time.Now().Nanosecond(),
	})

	token, err := loginToken.SignedString([]byte("SecretMostafa"))
	r, err := http.NewRequest("POST", url + "/request", bytes.NewReader([]byte("")))
	r.Header["Authorization"] = []string{"Bearer " + token}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("Status should be %d but it is %d", http.StatusForbidden, resp.StatusCode)
	}
}

func TestEndPointBWithRoleUserB(t *testing.T) {
	loginToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"role": "UserB",
		"exp":  time.Now().UTC().Add(time.Hour * 24).Unix(),
		"nano": time.Now().Nanosecond(),
	})

	token, err := loginToken.SignedString([]byte("SecretMostafa"))
	r, err := http.NewRequest("POST", url + "/request", bytes.NewReader([]byte("")))
	r.Header["Authorization"] = []string{"Bearer " + token}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Status should be %d but it is %d", http.StatusOK, resp.StatusCode)
	}
}

func TestEndPointBWithNonExistRole(t *testing.T) {
	loginToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"role": "NoRole",
		"exp":  time.Now().UTC().Add(time.Hour * 24).Unix(),
		"nano": time.Now().Nanosecond(),
	})

	token, err := loginToken.SignedString([]byte("SecretMostafa"))
	r, err := http.NewRequest("POST", url + "/request", bytes.NewReader([]byte("")))
	r.Header["Authorization"] = []string{"Bearer " + token}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusForbidden {
		t.Fatalf("Status should be %d but it is %d", http.StatusForbidden, resp.StatusCode)
	}
}
