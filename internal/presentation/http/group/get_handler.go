package httpgroup

import (
	"errors"
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetHandler struct {
	getGroupUseCase *group.GetGroup
	mapper          GroupResponseMapper
}

func NewGetHandler(uc *group.GetGroup, mapper GroupResponseMapper) *GetHandler {
	return &GetHandler{
		getGroupUseCase: uc,
		mapper:          mapper,
	}
}

func (h *GetHandler) Register(r gin.IRouter) {
	r.GET("/groups/:id", h.get)
}

func (h *GetHandler) get(c *gin.Context) {
	id := c.Param("id")
	g, err := h.getGroupUseCase.Run(id)
	if err != nil {
		if errors.Is(err, group.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gr := h.mapper.ToGroupResponse(g)
	c.JSON(http.StatusOK, gr)

}
