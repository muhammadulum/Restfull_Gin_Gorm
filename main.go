package main
import "github.com/gin-gonic/gin"
func main() {

	gin.SetMode(gin.ReleaseMode) 
	app := gin.Default()

	route := app
	route.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	route.Run(":8080")
}