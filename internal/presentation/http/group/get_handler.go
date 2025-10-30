package httpgroup

import (
	"kennen/internal/usecase"

	"github.com/gin-gonic/gin"
)

type GetHandler struct{ uc *usecase.GetGroup }

func NewGetHandler(uc *usecase.GetGroup) *GetHandler { return &GetHandler{uc: uc} }

func (h *GetHandler) Register(r *gin.Engine) {
	r.GET("/groups/:id", h.getGet)
}

func (h *GetHandler) getGet(c *gin.Context) {
	h.uc.Run()
}
