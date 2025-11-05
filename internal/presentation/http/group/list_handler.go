package httpgroup

import (
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListHandler struct{ listGroupsUseCase *group.ListGroup }

func NewListHandler(uc *group.ListGroup) *ListHandler {
	return &ListHandler{listGroupsUseCase: uc}
}

func (h *ListHandler) Register(r gin.IRouter) {
	r.GET("/groups", h.list)
}

func (h *ListHandler) list(c *gin.Context) {
	groups, err := h.listGroupsUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := make([]GroupResponse, 0, len(groups))
	for _, g := range groups {
		responses = append(responses, toGroupResponse(g))
	}
	c.JSON(http.StatusOK, responses)

}
