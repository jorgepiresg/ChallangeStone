package transfers

import (
	"context"
	"net/http"
	"time"

	"github.com/jorgepiresg/ChallangeStone/api/middleware"
	"github.com/jorgepiresg/ChallangeStone/app"
	model_transfer "github.com/jorgepiresg/ChallangeStone/model/transfer"
	"github.com/jorgepiresg/ChallangeStone/utils"
	"github.com/labstack/echo/v4"
)

type handler struct {
	app app.App
}

func Register(g *echo.Group, app app.App) {
	h := handler{
		app: app,
	}

	g.GET("", h.get, middleware.Private)
	g.POST("", h.do, middleware.Private)
}

// get godoc
// @Summary Get transfers
// @Description Gets the list of transfers from the authenticated user.
// @Tags         Transfers
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Token" default(Bearer <token>)
// @Success      200  {object}  model_transfer.Transfer
// @Failure      400  {object}  utils.Error
// @Router       /transfers [get]
func (h handler) get(c echo.Context) error {

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	accountOriginID := middleware.GetIDFromToken(c)

	tranfers, err := h.app.Transfer.GetByAccountID(ctx, accountOriginID)
	if err != nil {
		return utils.NewError(http.StatusBadRequest, "TODO", err.Error())
	}

	c.JSON(http.StatusOK, tranfers)
	return nil
}

// get godoc
// @Summary Do Transfers
// @Description transfers from one Account to another.
// @Tags         Transfers
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Token" default(Bearer <token>)
// @Param request body model_transfer.DoTransferRequest true "input"
// @Success      201
// @Failure      400  {object}  utils.Error
// @Router       /transfers [post]
func (h handler) do(c echo.Context) error {

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	var payload model_transfer.DoTransferRequest

	if err := c.Bind(&payload); err != nil {
		return utils.NewError(http.StatusBadRequest, "payload invalid ", err.Error())
	}

	if err := c.Validate(payload); err != nil {
		return utils.NewError(http.StatusBadRequest, "payload invalid, Missing params", err.Error())
	}

	accountOriginID := middleware.GetIDFromToken(c)
	accountDestinationID := payload.AccountDestinationID

	data := model_transfer.DoTransfer{
		AccountOriginID:      accountOriginID,
		AccountDestinationID: accountDestinationID,
		Amount:               payload.Amount,
	}

	err := h.app.Transfer.Do(ctx, data)
	if err != nil {
		return utils.NewError(http.StatusBadRequest, err.Error(), nil)
	}

	c.JSON(http.StatusCreated, nil)

	return nil
}
