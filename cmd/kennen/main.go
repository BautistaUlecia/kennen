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

	gr := infragroup.NewInMemoryRepository()
	rc := riot.NewClient(http.DefaultClient, apiKey)

	lguc := group.NewListGroup(gr)
	cguc := group.NewCreateGroup(gr)
	gguc := group.NewGetGroup(gr)
	atguc := group.NewAddToGroup(gr, rc)

	httpgroup.NewListHandler(lguc).Register(g)
	httpgroup.NewGetHandler(gguc).Register(g)
	httpgroup.NewCreateHandler(cguc).Register(g)
	httpgroup.NewAddHandler(atguc).Register(g)

	g.Run(":8080")
}
