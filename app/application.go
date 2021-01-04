package app

import ("github.com/gin-gonic/gin"
		"../logger"		
)

var(
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Log.Info("about to start the application")
	router.Run(":8080")
}