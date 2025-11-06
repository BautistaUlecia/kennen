package main

import (
	"fmt"
	infragroup "kennen/internal/infrastructure/group"
	"kennen/internal/infrastructure/riot"
	httpgroup "kennen/internal/presentation/http/group"
	"kennen/internal/routine"
	"kennen/internal/usecase/group"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("yes yes yes")
	_ = godotenv.Load()
	apiKey := os.Getenv("RIOT_API_KEY")

	// Esto creo que va en otro file tipo bootstrap.go
	g := gin.Default()

	// Configure CORS
	g.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200", "https://rivalry-ranks.com"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type"},
	}))

	// Create API group with /api prefix
	api := g.Group("/api")

	api.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	gr := infragroup.NewInMemoryRepository()
	rc := riot.NewClient(http.DefaultClient, apiKey)

	// Start version manager routine
	versionManager := routine.NewVersionManager(http.DefaultClient)
	versionManager.Start()

	// Create mapper with version manager
	mapper := httpgroup.NewGroupResponseMapper(versionManager)

	lguc := group.NewListGroup(gr)
	cguc := group.NewCreateGroup(gr)
	gguc := group.NewGetGroup(gr)
	atguc := group.NewAddToGroup(gr, rc)

	httpgroup.NewListHandler(lguc, mapper).Register(api)
	httpgroup.NewGetHandler(gguc, mapper).Register(api)
	httpgroup.NewCreateHandler(cguc, mapper).Register(api)
	httpgroup.NewAddHandler(atguc, mapper).Register(api)

	g.Run(":8080")
}
