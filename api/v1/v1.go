package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/jorgepiresg/ChallangeStone/api/middleware"
	"github.com/jorgepiresg/ChallangeStone/api/v1/account"
	"github.com/jorgepiresg/ChallangeStone/api/v1/transfers"
	"github.com/jorgepiresg/ChallangeStone/app"
	model_login "github.com/jorgepiresg/ChallangeStone/model/login"
	"github.com/jorgepiresg/ChallangeStone/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	app app.App
}

func Register(e *echo.Group, app app.App) {
	h := handler{
		app: app,
	}

	v1 := e.Group("/v1")

	v1.POST("/login", h.login)

	account.Register(v1.Group("/account"), app)
	transfers.Register(v1.Group("/transfers"), app)
}

// login godoc
// @Summary Login
// @Description login user to generate token
// @Tags         Login
// @Accept       json
// @Produce      json
// @Param request body model_login.Login true "input"
// @Success      200  {object}  model_login.LoginResponse
// @Failure      400  {object}  utils.Error
// @Router       /login [post]
func (h handler) login(c echo.Context) error {

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	var payload model_login.Login

	if err := c.Bind(&payload); err != nil {
		return utils.NewError(http.StatusBadRequest, "payload invalid ", err.Error())
	}

	payload.CPF = utils.CleanCPF(payload.CPF)

	account, err := h.app.Account.GetByCPF(ctx, payload.CPF)
	if err != nil && account.ID == "" {
		return utils.NewError(http.StatusBadRequest, "not found", err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(payload.Secret))
	if err != nil {
		return utils.NewError(http.StatusBadRequest, "not found", err.Error())
	}

	token, err := middleware.CreateJwtToken(account.ID)
	if err != nil {
		return utils.NewError(http.StatusBadRequest, "something went worng", err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"id":    "You were logged in!",
		"token": token,
	})
}
