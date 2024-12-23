package routes

import (
    "mvc/src/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
   
    router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testando",
		})
	})

    api := router.Group("/api")
    api.POST("/login", controllers.AuthLogin)

    api.POST("/users", controllers.AuthLogin)
    
	/*api.POST("/users", controllers.CreateUser)
    api.Use(middlewares.AuthMiddleware())
    {
        api.GET("/current-key", controllers.GetCurrentKey)
        api.GET("/keys", controllers.ListKeys)

        api.GET("/users", controllers.ListUsers) 
        api.PUT("/users/:id", controllers.UpdateUser) 
        api.DELETE("/users/:id", controllers.DeleteUser) 

        api.POST("/decrypt", controllers.DecryptDocument) 
    }*/

    return router
}
