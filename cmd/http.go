package cmd

import (
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/api"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	healthcheckAPI := &api.HealthCheckAPI{}
	e := echo.New()
	e.GET("/healthcheck",healthcheckAPI.HealthCheck)

	e.Start(":" + helpers.GetEnv("PORT", "9000"))
}