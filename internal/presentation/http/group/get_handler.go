package httpgroup

import (
	"errors"
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetHandler struct{ uc *group.GetGroup }

func NewGetHandler(uc *group.GetGroup) *GetHandler { return &GetHandler{uc: uc} }

func (h *GetHandler) Register(r *gin.Engine) {
	r.GET("/groups/:id", h.get)
}

func (h *GetHandler) get(c *gin.Context) {
	id := c.Param("id")
	g, err := h.uc.Run(id)
	if err != nil {
		if errors.Is(err, group.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, g)

}
