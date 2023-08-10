package account

import (
	"context"
	"net/http"
	"time"

	"github.com/jorgepiresg/ChallangeStone/app"
	model_account "github.com/jorgepiresg/ChallangeStone/model/account"
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

	g.GET("", h.list)
	g.GET("/:account_id/balance", h.getBalanceByAccountID)
	g.POST("", h.create)
}

// list godoc
// @Summary List of accounts
// @Description get list of accounts.
// @Tags Account
// @Produce      json
// @Success      200  {array}  model_account.Account
// @Failure      400  {object}  utils.Error
// @Router /account [get]
func (h handler) list(c echo.Context) error {

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	res, err := h.app.Account.List(ctx)
	if err != nil {
		return utils.NewError(http.StatusBadRequest, err.Error(), nil)
	}

	c.JSON(http.StatusOK, res)
	return nil
}

// getBalanceByAccountID godoc
// @Summary Account balance
// @Description get account balance
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        account_id   path      string  true  "Account ID"
// @Success      200  {object}  model_account.BalanceResponse
// @Failure      400  {object}  utils.Error
// @Router       /account/{account_id}/balance [get]
func (h handler) getBalanceByAccountID(c echo.Context) error {

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	id := c.Param("account_id")

	balance, err := h.app.Account.BalanceByAccountID(ctx, id)
	if err != nil {
		return utils.NewError(http.StatusBadRequest, err.Error(), nil)
	}

	res := model_account.BalanceResponse{
		Balance: balance,
	}

	c.JSON(http.StatusOK, res)

	return nil
}

// create godoc
// @Summary Account create
// @Description create a account
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param request body model_account.Create true "input"
// @Success      201  {object}  model_account.CreateResponse
// @Failure      400  {object}  utils.Error
// @Router       /account [post]
func (h handler) create(c echo.Context) error {

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	var payload model_account.Create

	if err := c.Bind(&payload); err != nil {
		return utils.NewError(http.StatusBadRequest, "payload invalid ", err.Error())
	}

	if err := c.Validate(payload); err != nil {
		return utils.NewError(http.StatusBadRequest, "payload invalid, Missing params", err.Error())
	}

	account, err := h.app.Account.Create(ctx, payload)
	if err != nil {
		return utils.NewError(http.StatusBadRequest, "fail to create a new account", err.Error())
	}

	res := model_account.CreateResponse{
		ID: account.ID,
	}

	c.JSON(http.StatusCreated, res)

	return nil
}
