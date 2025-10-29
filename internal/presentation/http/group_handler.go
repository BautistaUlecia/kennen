package httpgroup

import (
	"kennen/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Create usecase.CreateGroup
	Show   usecase.ShowGroup
}

func (h *Handler) Register(r *gin.Engine) {
	r.GET("/groups/:id", h.showGroup)
	r.POST("/groups", h.createGroup)
}

func New(create usecase.CreateGroup, show usecase.ShowGroup) *Handler {
	return &Handler{Create: create, Show: show}
}

func (h *Handler) createGroup(c *gin.Context) {}
func (h *Handler) showGroup(c *gin.Context) {
	h.Show.Run()
}
