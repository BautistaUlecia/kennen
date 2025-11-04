package httpgroup

import (
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddHandler struct {
	addToGroupUseCase *group.AddToGroup
}
type AddToGroupRequest struct {
	Region   string `json:"region"`
	GameName string `json:"game_name"`
	Tag      string `json:"tag"`
}

func NewAddHandler(uc *group.AddToGroup) *AddHandler {
	return &AddHandler{addToGroupUseCase: uc}
}

func (h *AddHandler) Register(r *gin.Engine) {
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
	err = h.addToGroupUseCase.Run(groupID, request.GameName, request.Tag, request.Region)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}
