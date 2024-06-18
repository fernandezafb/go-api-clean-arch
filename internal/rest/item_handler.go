package rest

import (
	"net/http"
	"strconv"

	"github.com/fernandezafb/go-api-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	ItemUsecase domain.ItemUsecase
}

func (i *ItemHandler) SellItemHandler(c echo.Context) error {
	request := domain.ItemRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	r, err := i.ItemUsecase.SellItem(request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, r)
}

func (i *ItemHandler) ShowItemPropertiesHandler(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	r, err := i.ItemUsecase.ShowItemProperties(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, r)
}

func (i *ItemHandler) UpdateItemPropertiesHandler(c echo.Context) error {
	request := domain.ItemRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	r, err := i.ItemUsecase.UpdateItemProperties(id, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, r)
}

func (i *ItemHandler) DeleteItemHandler(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	r, err := i.ItemUsecase.DeleteItem(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, r)
}
