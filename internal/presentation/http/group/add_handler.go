package httpgroup

import (
	"kennen/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AddHandler struct {
	uc *usecase.AddToGroup
}

func NewAddHandler(uc *usecase.AddToGroup) *AddHandler {
	return &AddHandler{uc: uc}
}

func (h *AddHandler) Register(r *gin.Engine) {
	r.POST("/groups/:id/summoners", h.postAddSummoner)
}

func (h *AddHandler) postAddSummoner(c *gin.Context) {
	h.uc.Run()
}
