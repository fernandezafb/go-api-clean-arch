package main

import (
	"log/slog"
	"os"

	"github.com/fernandezafb/go-api-clean-arch/domain"
	"github.com/fernandezafb/go-api-clean-arch/internal/rest"
	"github.com/fernandezafb/go-api-clean-arch/internal/storage/inmemory"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Specification struct {
	Port    string `default:"9450"`
	Version string `required:"true"`
}

// TODO: 1. Extract config environment logic. 2. Configure custom echo logger to use slog. 3. Add a custom auth middleware.
func main() {
	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Load env vars
	var s Specification
	err := envconfig.Process("", &s)
	if err != nil {
		slog.Error("Failed to process env var", "error", err)
		os.Exit(1)
	}

	e := echo.New()

	// Custom error handler
	e.HTTPErrorHandler = rest.NewHTTPErrorHandler(domain.CustomErrors()).CustomHandler

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	rest.BootstrapRouter(inmemory.NewDb(), e)

	// Start server
	e.Logger.Fatal(e.Start(":9450"))
}
