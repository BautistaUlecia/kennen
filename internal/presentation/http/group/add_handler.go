package httpgroup

import (
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddHandler struct {
	uc *group.AddToGroup
}
type AddToGroupRequest struct {
	Region   string
	GameName string
	Tag      string
}

func NewAddHandler(uc *group.AddToGroup) *AddHandler {
	return &AddHandler{uc: uc}
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
	err = h.uc.Run(groupID, request.GameName, request.Tag, request.Region)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}
