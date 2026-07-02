package routes

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	api := router.Group("/api")
	{
		RegisterStudentRoutes(api)
		// RegisterSubjectRoutes(api)
	}
}
