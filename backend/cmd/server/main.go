package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kenstillings/ticks-and-tomes/internal/config"
	"github.com/kenstillings/ticks-and-tomes/internal/handlers"
)

func main() {
	// Load environment
	cfg := config.NewConfig()

	// Initialize router
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Routes
	api := router.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		auth.POST("/login", handlers.Login)
		auth.POST("/register", handlers.Register)
		auth.POST("/logout", handlers.Logout)

		// Empire routes
		empire := api.Group("/empire")
		empire.GET("", handlers.GetEmpire)
		empire.POST("", handlers.CreateEmpire)
		empire.PUT("", handlers.UpdateEmpire)

		// Action routes
		action := api.Group("/action")
		action.POST("/explore", handlers.ActionExplore)
		action.POST("/meditate", handlers.ActionMeditate)
		action.POST("/drill", handlers.ActionDrill)
		action.POST("/farm", handlers.ActionFarm)

		// Spell routes
		spell := api.Group("/spell")
		spell.POST("/love", handlers.CastLoveSpell)
		spell.POST("/shield", handlers.CastShieldSpell)

		// Chat routes
		chat := api.Group("/chat")
		chat.GET("/messages", handlers.GetMessages)
		chat.POST("/messages", handlers.SendMessage)

		// Market routes
		market := api.Group("/market")
		market.GET("/listings", handlers.GetListings)
		market.POST("/trade", handlers.PlaceTrade)

		// Clan routes
		clan := api.Group("/clan")
		clan.GET("", handlers.GetClan)
		clan.POST("", handlers.CreateClan)
		clan.PUT("/:id", handlers.UpdateClan)
	}

	// Start server
	port := cfg.BackendPort
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
