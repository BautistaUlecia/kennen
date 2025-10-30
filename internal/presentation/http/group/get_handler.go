package httpgroup

import (
	"kennen/internal/usecase/group"

	"github.com/gin-gonic/gin"
)

type GetHandler struct{ uc *group.GetGroup }

func NewGetHandler(uc *group.GetGroup) *GetHandler { return &GetHandler{uc: uc} }

func (h *GetHandler) Register(r *gin.Engine) {
	r.GET("/groups/:id", h.getGet)
}

func (h *GetHandler) getGet(c *gin.Context) {
	h.uc.Run()
}
