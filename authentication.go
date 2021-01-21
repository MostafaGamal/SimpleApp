package main

import (
	"fmt"
	"net/http"
	"time"

	"WeVentureTask/ent"
	"WeVentureTask/ent/users"
	"github.com/alexedwards/argon2id"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func (app *app) Login(c echo.Context) error {
	var m echo.Map
	if err := c.Bind(&m); err != nil {
		return err
	}

	if val, ok := m["username"]; !ok || val == "" {
		return c.String(http.StatusBadRequest, "invalid username")
	}

	if val, ok := m["password"]; !ok || val == "" {
		return c.String(http.StatusBadRequest, "invalid password")
	}

	u, err := app.dbCtx.Client.Users.
		Query().
		Where(users.UsernameEQ(m["username"].(string))).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(app.dbCtx.Ctx)

	if err != nil {
		switch err.(type) {
		case *ent.NotFoundError:
			return c.String(http.StatusUnauthorized, "invalid username/password")
		case *ent.NotSingularError:
			return c.String(http.StatusUnauthorized, "invalid username/password")
		default:
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	match, err := argon2id.ComparePasswordAndHash(m["password"].(string), u.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal error occurred")
	}

	if !match {
		return c.String(http.StatusUnauthorized, "invalid username/password")
	}

	loginToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"role": u.Role,
		"exp":  time.Now().UTC().Add(time.Hour * 24).Unix(),
		"nano": time.Now().Nanosecond(),
	})

	token, err := loginToken.SignedString([]byte("SecretMostafa"))
	if err != nil {
		return fmt.Errorf("failed querying user: %v", err)
	}

	return c.JSON(http.StatusOK, map[string]string{"token" : token})
}
