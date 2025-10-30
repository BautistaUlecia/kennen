package main

import (
	"fmt"
	httpgroup "kennen/internal/presentation/http/group"
	"kennen/internal/usecase/group"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("yes yes yes")
	_ = godotenv.Load()

	// Esto creo que va en otro file tipo bootstrap.go
	g := gin.Default()
	g.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	createGroupUseCase := group.NewCreateGroup()
	getGroupUseCase := group.NewGetGroup()
	addToGroupUseCase := group.NewAddToGroup()

	httpgroup.NewGetHandler(getGroupUseCase).Register(g)
	httpgroup.NewCreateHandler(createGroupUseCase).Register(g)
	httpgroup.NewAddHandler(addToGroupUseCase).Register(g)

	g.Run(":8080")
}
