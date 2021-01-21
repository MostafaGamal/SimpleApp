package main

import (
	"errors"
	"strings"

	"WeVentureTask/casbinrest"
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/gommon/log"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

const (
	BaseURL       = "/app"
	LoginPath     = "/login"
	authModelPath = "auth_model.conf"
	policyPath    = "policy.csv"
)

// Our main app struct
type app struct {
	e     *echo.Echo
	dbCtx *DBCtx
}

// GetRoleByToken implements casbinrest adapter DataSource Interface
func (app *app) GetRoleByToken(reqToken string) string {
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != "HS512" ||
			token.Claims.(jwt.MapClaims).Valid() != nil {
			return nil, errors.New("invalid/expired token")
		}

		return []byte("SecretMostafa"), nil
	})

	if err != nil {
		return ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}

	role, ok := claims["role"].(string)
	if !ok {
		return ""
	}

	return role
}

// NewApp creates a concrete AppInterface implementation
func NewApp() AppInterface {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	ce, err := casbin.NewEnforcer(authModelPath, policyPath)
	if err != nil {
		e.Logger.Fatal(err)
	}

	dbCtx, err := NewDBClient()
	if err != nil {
		e.Logger.Fatal(err)
	}

	server := &app{
		e:     e,
		dbCtx: dbCtx,
	}

	group := e.Group(BaseURL)
	group.Use(casbinrest.MiddlewareWithConfig(casbinrest.Config{
		Skipper: func(ctx echo.Context) bool {
			// Skip authentication for login request
			return ignoreAuthentication(ctx)
		},
		Enforcer: ce,
		Source:   server,
	}))

	RegisterRoutes(group, server)

	return server
}

func ignoreAuthentication(ctx echo.Context) bool {
	return ctx.Path() == BaseURL+LoginPath &&
		ctx.Request().Method == "POST"
}

func main() {
	var newApp *app
	application := NewApp()

	switch application.(type) {
	case *app:
		newApp = application.(*app)
	default:
		panic("error creating app")
	}

	defer newApp.dbCtx.Client.Close()
	err := newApp.dbCtx.CreateSchema()
	if err != nil {
		newApp.e.Logger.Fatalf("failed creating schema resources: %v", err)
	}

	_, err = newApp.dbCtx.CreateUsers()
	if err != nil && !strings.Contains(err.Error(), "users already exist") {
		newApp.e.Logger.Fatal(err)
	}

	newApp.e.Logger.Fatal(newApp.e.Start(":8080"))
}
