package main

import (
	"github.com/AKiRA3563/sa-65/controller"

	"github.com/AKiRA3563/sa-65/entity"

	"github.com/AKiRA3563/sa-65/middlewares"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			r.POST("/createborrowlist", controller.CreateBorrowList)
			r.GET("/borrowlist", controller.ListBorrow)
			
			r.GET("/equipment/:id", controller.GetEquipment)
			r.GET("/equipment", controller.ListEquipment)

			r.GET("/user/:id", controller.GetUser)
			r.GET("/user", controller.ListUser)
			
			r.GET("/employee/:id", controller.GetEmployee)
			r.GET("/employee", controller.ListEmployee)
		}
	}

	// login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
