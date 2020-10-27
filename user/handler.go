package user

import (
	"errors"
	"github.com/Leonardo-Antonio/api-notes/helper"
	"github.com/labstack/echo"
	"net/http"
)

type handlerU struct {
	storage iuser
}

func newHandler(s iuser) *handlerU {
	return &handlerU{s}
}

func (h *handlerU) SignUp(c echo.Context) error {
	user := model{}
	err := c.Bind(&user)
	if err != nil {
		response := helper.NewResponseJSON(
			helper.ERROR,
			"the entered structure is not valid",
			true,
			nil,
			)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = h.storage.CreateUser(user)
	if err != nil {
		if errors.Is(err, helper.ErrEmailInvalid) ||
			errors.Is(err, helper.ErrPasswordNotSecure) ||
			errors.Is(err, helper.ErrNickNameIvalid) ||
			errors.Is(err, helper.ErrDataInvalid) ||
			errors.Is(err, helper.ErrInsertionFailed){
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}
	response := helper.NewResponseJSON(
		helper.MESSAGE,
		"the account was created successfully",
		false,
		nil,
		)
	return c.JSON(http.StatusCreated, response)
}

func (h *handlerU) LogIn(c echo.Context) error {
	credentials := model{}
	err := c.Bind(&credentials)
	if err != nil {
		response := helper.NewResponseJSON(
			helper.ERROR,
			"the entered structure is not valid",
			true,
			nil,
		)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := h.storage.Login(credentials)
	if err != nil {
		if errors.Is(err, helper.ErrEmailInvalid) ||
			errors.Is(err, helper.ErrPasswordNotSecure) ||
			errors.Is(err, helper.ErrDataInvalid) {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helper.NewResponseJSON(helper.ERROR, "OK", false, data)
	return c.JSON(http.StatusOK, response)
}