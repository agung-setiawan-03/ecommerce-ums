package api

import (
	"ecommerce-ums/constants"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserAPI struct {
	UserService interfaces.IUserService
}

func (api *UserAPI) RegisterUser(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	req := models.User{}

	if err := e.Bind(&req); err != nil {
		log.Error("Failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	resp, err := api.UserService.RegisterUser(e.Request().Context(), &req)
	if err != nil {
		log.Error("Failed to register user: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *UserAPI) RegisterAdmin(e echo.Context) error {
	var (
		log = helpers.Logger
	)
	req := models.User{}

	if err := e.Bind(&req); err != nil {
		log.Error("Failed to parse request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request: ", err)
		return helpers.SendResponseHTTP(e, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
	}

	resp, err := api.UserService.RegisterAdmin(e.Request().Context(), &req)
	if err != nil {
		log.Error("Failed to register user: ", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)
	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}
