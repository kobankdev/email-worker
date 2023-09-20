package app

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getEmailRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) getEmail(ctx *gin.Context) {
	var req getEmailRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	user, err := server.store.GetEmail(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
