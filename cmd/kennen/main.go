package main

import (
	"fmt"
	infragroup "kennen/internal/infrastructure/group"
	"kennen/internal/infrastructure/riot"
	httpgroup "kennen/internal/presentation/http/group"
	"kennen/internal/usecase/group"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("yes yes yes")
	_ = godotenv.Load()
	apiKey := os.Getenv("RIOT_API_KEY")

	// Esto creo que va en otro file tipo bootstrap.go
	g := gin.Default()
	g.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	groupRepository := infragroup.NewInMemoryRepository()
	riotClient := riot.NewClient(http.DefaultClient, apiKey)

	listGroupUseCase := group.NewListGroup(groupRepository)
	createGroupUseCase := group.NewCreateGroup(groupRepository)
	getGroupUseCase := group.NewGetGroup(groupRepository)
	addToGroupUseCase := group.NewAddToGroup(groupRepository, riotClient)

	httpgroup.NewListHandler(listGroupUseCase).Register(g)
	httpgroup.NewGetHandler(getGroupUseCase).Register(g)
	httpgroup.NewCreateHandler(createGroupUseCase).Register(g)
	httpgroup.NewAddHandler(addToGroupUseCase).Register(g)

	g.Run(":8080")
}
