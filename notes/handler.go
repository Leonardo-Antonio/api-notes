package notes

import (
	"errors"
	"github.com/Leonardo-Antonio/api-notes/helper"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type handNote struct {
	storage iNotes
}

func newHandNote(s iNotes) *handNote {
	return &handNote{s}
}

func (n *handNote) NewNote(c echo.Context) error {
	note := model{}
	err := c.Bind(&note)
	if err != nil {
		response := helper.NewResponseJSON(helper.ERROR, "the structure is not valid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = n.storage.NewNote(note)
	if err != nil {
		if errors.Is(err, helper.ErrIDInvalid) ||
			errors.Is(err, helper.ErrDataInvalid) ||
			errors.Is(err, helper.ErrInsertionFailed) {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}
	response := helper.NewResponseJSON(helper.MESSAGE, "note created successfully", false, nil)
	return c.JSON(http.StatusCreated, response)
}

func (n *handNote) UpdateNote(c echo.Context) error {
	note := model{}
	err := c.Bind(&note)
	if err != nil {
		response := helper.NewResponseJSON(helper.ERROR, "the structure is not valid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = n.storage.UpdateNote(note)
	if err != nil {
		if errors.Is(err, helper.ErrIDInvalid) ||
			errors.Is(err, helper.ErrDataInvalid) ||
			errors.Is(err, helper.ErrRowNotAffected) {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}
	response := helper.NewResponseJSON(helper.MESSAGE, "note updated successfully", false, nil)
	return c.JSON(http.StatusOK, response)
}

func (n *handNote) DeleteNote(c echo.Context) error {
	note := model{}
	err := c.Bind(&note)
	if err != nil {
		response := helper.NewResponseJSON(helper.ERROR, "the structure is not valid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = n.storage.DeleteNote(note)
	if err != nil {
		if errors.Is(err, helper.ErrIDInvalid) ||
			errors.Is(err, helper.ErrDataInvalid) ||
			errors.Is(err, helper.ErrRowNotAffected) {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}
	response := helper.NewResponseJSON(helper.MESSAGE, "note deleted successfully", false, nil)
	return c.JSON(http.StatusOK, response)
}

func (n *handNote) GetAllNote(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.NewResponseJSON(helper.ERROR, "the parameter must be an integer", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data, err := n.storage.GetAllNotes(ID)
	if err != nil {
		response := helper.NewResponseJSON(helper.ERROR, err.Error(), true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := helper.NewResponseJSON(helper.MESSAGE, "OK", false, data)
	return c.JSON(http.StatusOK, response)
}
