package httpgroup

import (
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddHandler struct {
	addToGroupUseCase *group.AddToGroup
	mapper            GroupResponseMapper
}

func NewAddHandler(uc *group.AddToGroup, mapper GroupResponseMapper) *AddHandler {
	return &AddHandler{
		addToGroupUseCase: uc,
		mapper:            mapper,
	}
}

func (h *AddHandler) Register(r gin.IRouter) {
	r.POST("/groups/:id/summoners", h.addToGroup)
}

func (h *AddHandler) addToGroup(c *gin.Context) {
	groupID := c.Param("id")
	var request AddToGroupRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	g, err := h.addToGroupUseCase.Run(groupID, request.GameName, request.Tag, request.Region)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	gr := h.mapper.ToGroupResponse(g)
	c.JSON(http.StatusOK, gr)
}
