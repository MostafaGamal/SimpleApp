package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"WeVentureTask/ent"
	"github.com/alexedwards/argon2id"
)

type DBCtx struct {
	Client *ent.Client
	Ctx    context.Context
}

// NewDBClient initialize ent DBClient
func NewDBClient() (dbContext *DBCtx, err error) {
	client, err := ent.Open("sqlite3", "file:weventure.db?_busy_timeout=300000&cache=shared&_fk=1")
	if err != nil {
		return
	}

	dbContext = &DBCtx{
		Client: client,
		Ctx:    context.Background(),
	}

	return
}

// CreateSchema Creates DB Schema
func (dbCtx *DBCtx) CreateSchema() (err error) {
	// Run the auto migration tool.
	err = dbCtx.Client.Schema.Create(dbCtx.Ctx)
	return
}

// CreateUsers Populate DB with users
func (dbCtx *DBCtx) CreateUsers() ([]*ent.Users, error) {
	password, err := argon2id.CreateHash("testtest", argon2id.DefaultParams)
	if err != nil {
		return nil, fmt.Errorf("failed hashing user password: %v", err)
	}

	password2, err := argon2id.CreateHash("testtest2", argon2id.DefaultParams)
	if err != nil {
		return nil, fmt.Errorf("failed hashing user password: %v", err)
	}

	users, err := dbCtx.Client.Users.CreateBulk(
		dbCtx.Client.Users.Create().SetName("Mostafa").
			SetUsername("mostafa").SetPassword(password).
			SetRole("UserA"),
		dbCtx.Client.Users.Create().SetName("Mai").
			SetUsername("mai").SetPassword(password2).
			SetRole("UserB"),
	).Save(dbCtx.Ctx)
	if err != nil {
		if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, fmt.Errorf("failed creating users: %v", err)
		}
		return nil, errors.New("users already exist")
	}
	return users, nil
}
