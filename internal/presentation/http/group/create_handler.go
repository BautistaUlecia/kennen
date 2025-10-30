package httpgroup

import (
	"kennen/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CreateHandler struct {
	uc *usecase.CreateGroup
}

func NewCreateHandler(uc *usecase.CreateGroup) *CreateHandler {
	return &CreateHandler{uc: uc}
}
func (h *CreateHandler) Register(r *gin.Engine) {
	r.POST("/groups", h.postCreate)
}

func (h *CreateHandler) postCreate(c *gin.Context) {
	h.uc.Run()

}
