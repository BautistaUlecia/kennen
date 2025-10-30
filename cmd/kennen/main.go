package main

import (
	"fmt"
	httpgroup "kennen/internal/presentation/http/group"
	"kennen/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("yes yes yes")
	_ = godotenv.Load()

	g := gin.Default()
	g.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	createGroupUseCase := usecase.NewCreateGroup()
	getGroupUseCase := usecase.NewGetGroup()
	addToGroupUseCase := usecase.NewAddToGroup()

	httpgroup.NewGetHandler(getGroupUseCase).Register(g)
	httpgroup.NewCreateHandler(createGroupUseCase).Register(g)
	httpgroup.NewAddHandler(addToGroupUseCase).Register(g)

	g.Run(":8080")
}
