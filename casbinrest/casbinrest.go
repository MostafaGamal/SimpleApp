package casbinrest

import (
	"strings"

	"github.com/casbin/casbin/v2"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	// Config defines the config for CasbinAuth middleware.
	Config struct {
		Skipper  middleware.Skipper
		Enforcer *casbin.Enforcer
		Source   DataSource
	}
)

var (
	// DefaultConfig is the default CasbinAuth middleware config.
	DefaultConfig = Config{
		Skipper: middleware.DefaultSkipper,
	}
)

// DataSource is the Authen from datasource
type DataSource interface {
	GetRoleByToken(reqToken string) string
}

// Middleware returns a CasbinAuth middleware.
func Middleware(ce *casbin.Enforcer, sc DataSource) echo.MiddlewareFunc {
	c := DefaultConfig
	c.Enforcer = ce
	c.Source = sc
	return MiddlewareWithConfig(c)
}

// MiddlewareWithConfig returns a CasbinAuth middleware with config.
func MiddlewareWithConfig(config Config) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultConfig.Skipper
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if ok, _ := config.CheckPermission(c); ok || config.Skipper(c) {
				return next(c)
			}
			return echo.ErrForbidden
		}
	}
}

// GetRole gets the role name from the request.
func (a *Config) GetRole(c echo.Context) string {
	reqToken := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return ""
	}
	return a.Source.GetRoleByToken(strings.TrimSpace(splitToken[1]))
}

// CheckPermission checks the role/path/method combination from the request.
func (a *Config) CheckPermission(c echo.Context) (bool, error) {
	return a.Enforcer.Enforce(a.GetRole(c), c.Request().URL.Path, c.Request().Method)
}
