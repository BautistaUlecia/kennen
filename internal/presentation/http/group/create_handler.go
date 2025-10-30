package httpgroup

import (
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGroupRequest struct {
	Name string
}

type CreateHandler struct {
	createGroupUseCase *group.CreateGroup
}

func NewCreateHandler(createGroupUseCase *group.CreateGroup) *CreateHandler {
	return &CreateHandler{createGroupUseCase: createGroupUseCase}
}
func (h *CreateHandler) Register(r *gin.Engine) {
	r.POST("/groups", h.create)
}

func (h *CreateHandler) create(c *gin.Context) {
	var request CreateGroupRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	h.createGroupUseCase.Run(request.Name)
}
