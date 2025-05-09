package main

import (
	"explore-aks-backend-app-air/src/cache"
	"explore-aks-backend-app-air/src/constants"
	"explore-aks-backend-app-air/src/database"
	"explore-aks-backend-app-air/src/routes"
	"explore-aks-backend-app-air/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	isRelease := constants.ENV_MODE == "release"

	// Initialize the Gin router
	router := gin.Default()

	// Connect to the database
	database.ConnectDB()

	// Connect to the cache
	cache.ConnectCache()

	// Setup routes
	routes.SetupRoutes(router)

	// Release actions
	if isRelease {
		logrus.Info("Running in release mode")
		gin.SetMode(gin.ReleaseMode)
	} else {
		logrus.Info("Running in debug mode")
		utils.ExecuteDebugingActions()
	}

	// Start the server
	router.Run(":8080")
}
