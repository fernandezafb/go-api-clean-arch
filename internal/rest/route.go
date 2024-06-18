package rest

import (
	"net/http"

	"github.com/fernandezafb/go-api-clean-arch/internal/storage/inmemory"
	"github.com/fernandezafb/go-api-clean-arch/internal/storage/repository"
	"github.com/fernandezafb/go-api-clean-arch/internal/usecase"
	"github.com/labstack/echo/v4"
)

func NewItemRouter(db inmemory.DbClient, e *echo.Echo) {
	ir := repository.NewItemRepository(db)
	ih := ItemHandler{
		ItemUsecase: usecase.NewItemUsecase(ir),
	}

	e.POST("/items", ih.SellItemHandler)
	e.GET("/items/:id", ih.ShowItemPropertiesHandler)
	e.PUT("/items/:id", ih.UpdateItemPropertiesHandler)
	e.DELETE("/items/:id", ih.DeleteItemHandler)
}

func BootstrapRouter(db inmemory.DbClient, e *echo.Echo) {
	// Public APIs
	e.GET("/", func(c echo.Context) error {
		status := struct {
			Status string `json:"status"`
		}{Status: "healthy"}

		return c.JSON(http.StatusOK, status)
	})

	NewItemRouter(db, e)
}
