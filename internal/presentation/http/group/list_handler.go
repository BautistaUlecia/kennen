package httpgroup

import (
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListHandler struct {
	listGroupsUseCase *group.ListGroup
	versionGetter     VersionGetter
}

func NewListHandler(uc *group.ListGroup, vg VersionGetter) *ListHandler {
	return &ListHandler{
		listGroupsUseCase: uc,
		versionGetter:     vg,
	}
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
		responses = append(responses, toGroupResponse(g, h.versionGetter))
	}
	c.JSON(http.StatusOK, responses)

}
