package type_notes

import (
	"errors"
	"github.com/Leonardo-Antonio/api-notes/helper"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type handlerTN struct {
	storage iTypeNote
}

func newHandlerTN(s iTypeNote) *handlerTN {
	return &handlerTN{s}
}

func (t *handlerTN) NewActivity(c echo.Context) error {
	TypeNote := model{}
	err := c.Bind(&TypeNote)
	if err != nil {
		response := helper.NewResponseJSON(helper.ERROR, "the entered structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = t.storage.CreateActivity(TypeNote.TypeOfNote)
	if err != nil {
		response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.NewResponseJSON(
		helper.MESSAGE,
		"successfully created activity",
		true,
		nil,
		)
	return c.JSON(http.StatusCreated, response)
}

func (t *handlerTN) GetTypes(c echo.Context) error {
	data, err := t.storage.GetTypes()
	if err != nil {
		response := helper.NewResponseJSON(helper.MESSAGE, err.Error(), true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.NewResponseJSON(helper.MESSAGE, "OK", false, data)
	return c.JSON(http.StatusOK, response)
}

func (t *handlerTN) DeleteType(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.NewResponseJSON(helper.ERROR, "the parameter must be a numeric data", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = t.storage.DeleteType(ID)
	if err != nil {
		if errors.Is(err, helper.ErrRowNotAffected) {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helper.NewResponseJSON(helper.MESSAGE, "was successfully removed", false, nil)
	return c.JSON(http.StatusOK, response)
}