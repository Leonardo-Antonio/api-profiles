package user

import (
	"errors"
	"github.com/Leonardo-Antonio/api-profiles/helper"
	"github.com/labstack/echo"
	"net/http"
)

type handler struct {
	storage iuser
}

func newHandler(s iuser) *handler {
	return &handler{s}
}

func (h *handler) GetAll(c echo.Context) error {
	users, err := h.storage.GetAll()
	if err != nil {
		response := helper.ResponseJSON(helper.ERROR, err.Error(), true, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := helper.ResponseJSON(helper.MESSAGE, "OK", false, users)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) SignUp(c echo.Context) error {
	user := model{}
	err := c.Bind(&user)
	if err != nil {
		response := helper.ResponseJSON(helper.ERROR, "the structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.storage.SignUp(user)
	if err != nil {
		if errors.Is(err, helper.ErrEmailInvalid) ||
			errors.Is(err, helper.ErrPersonalInformationInvalid) ||
			errors.Is(err, helper.ErrInsecurePassword) ||
			errors.Is(err, helper.ErrRowNotAffected) {
			response := helper.ResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		}
		response := helper.ResponseJSON(helper.ERROR, err.Error(), true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseJSON(helper.MESSAGE, "user created successfully", false, nil)
	return c.JSON(http.StatusCreated, response)
}

func (h *handler) SignIn(c echo.Context) error {
	identification := model{}
	err := c.Bind(&identification)

	if err != nil {
		response := helper.ResponseJSON(helper.ERROR, "the structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := h.storage.SignIn(identification)
	if err != nil {
		if errors.Is(err, helper.ErrStmtSQL) {
			response := helper.ResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		} else {
			response := helper.ResponseJSON(helper.ERROR, "the user not exits", true, nil)
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	response := helper.ResponseJSON(helper.MESSAGE, "OK", false, data)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) Delete(c echo.Context) error {
	identification := model{}
	err := c.Bind(&identification)
	if err != nil {
		response := helper.ResponseJSON(helper.ERROR, "the structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.storage.Delete(identification)
	if err != nil {
		if errors.Is(err, helper.ErrUserInvalid) {
			response := helper.ResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		} else {
			response := helper.ResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response := helper.ResponseJSON(helper.MESSAGE, "user successfully removed", false, nil)
	return c.JSON(http.StatusOK, response)
}

func (h *handler) Update(c echo.Context) error {
	data := model{}
	err := c.Bind(&data)
	if err != nil {
		response := helper.ResponseJSON(helper.ERROR, "the structure is invalid", true, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.storage.Update(data)
	if err != nil {
		if errors.Is(err, helper.ErrStmtSQL) {
			response := helper.ResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusInternalServerError, response)
		} else {
			response := helper.ResponseJSON(helper.ERROR, err.Error(), true, nil)
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	response := helper.ResponseJSON(helper.MESSAGE, "user updated successfully", false, nil)
	return c.JSON(http.StatusOK, response)
}
