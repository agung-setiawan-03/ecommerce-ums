package cmd

import (
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/api"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/repository"
	"ecommerce-ums/internal/services"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	d := dependencyInject()

	e := echo.New()
	e.GET("/healthcheck", d.HealtcheckAPI.HealthCheck)

	userV1 := e.Group("/user/v1")
	userV1.POST("/register", d.UserAPi.RegisterUser)
	userV1.POST("/register/admin", d.UserAPi.RegisterAdmin)

	e.Start(":" + helpers.GetEnv("PORT", "9000"))
}

type Dependency struct {
	HealtcheckAPI *api.HealthCheckAPI
	UserAPi       interfaces.IUserAPI
}

func dependencyInject() Dependency {
	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}
	userSvc := &services.UserService{
		UserRepo: userRepo,
	}
	userAPI := &api.UserAPI{
		UserService: userSvc,
	}
	return Dependency{
		HealtcheckAPI: &api.HealthCheckAPI{},
		UserAPi:       userAPI,
	}
}
