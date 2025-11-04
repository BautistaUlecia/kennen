package httpgroup

import (
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListHandler struct{ uc *group.ListGroup }

func NewListHandler(uc *group.ListGroup) *ListHandler { return &ListHandler{uc: uc} }

func (h *ListHandler) Register(r *gin.Engine) {
	r.GET("/groups", h.list)
}

func (h *ListHandler) list(c *gin.Context) {
	g, err := h.uc.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, g)

}
