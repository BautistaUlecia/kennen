package httpgroup

import (
	"errors"
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateHandler struct {
	createGroupUseCase *group.CreateGroup
	mapper             GroupResponseMapper
}

func NewCreateHandler(uc *group.CreateGroup, mapper GroupResponseMapper) *CreateHandler {
	return &CreateHandler{
		createGroupUseCase: uc,
		mapper:             mapper,
	}
}

func (h *CreateHandler) Register(r gin.IRouter) {
	r.POST("/groups", h.create)
}

func (h *CreateHandler) create(c *gin.Context) {
	var request CreateGroupRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	g, err := h.createGroupUseCase.Run(request.Name)
	if err != nil {
		if errors.Is(err, group.ErrNameTaken) {
			c.JSON(http.StatusConflict, gin.H{"error": "group name already taken"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	gr := h.mapper.ToGroupResponse(g)
	c.JSON(http.StatusCreated, gr)
}
