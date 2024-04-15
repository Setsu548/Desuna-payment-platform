package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/db/util"
	"github.com/gin-gonic/gin"
)

// loginUserRequest defines the request body for loginUser API
type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

// loginUserResponse defines the response body for loginUser API
type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (q *ServerConfig) LoginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user, err := q.Store.GetUserByName(ctx, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
		return
	}

	accessToken, err := q.TokenMaker.CreateToken(user.Username, q.Config.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	resp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, resp)
}
