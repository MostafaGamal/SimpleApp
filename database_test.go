package main_test

import (
	"fmt"
	"strings"
	"testing"

	main "WeVentureTask"
	"WeVentureTask/ent"
)

func TestCreateUsers(t *testing.T) {
	dbCtx, err := main.NewDBClient()
	if err != nil {
		t.Fatal(err)
	}

	createdUsers, err := dbCtx.CreateUsers()
	if err != nil && !strings.Contains(err.Error(), "users already exist") {
		t.Fatal(err)
	}

	dbUsers, err := QueryUsers(dbCtx)
	if err != nil {
		t.Fatal(err)
	}

	if createdUsers != nil && len(createdUsers) != len(dbUsers) {
		t.Fatal("error creating users")
	}

	if len(dbUsers) != 2 {
		t.Fatal("invalid db users count")
	}
}

func QueryUsers(dbCtx *main.DBCtx) ([]*ent.Users, error) {
	dbUsers, err := dbCtx.Client.Users.
		Query().All(dbCtx.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}

	return dbUsers, nil
}
