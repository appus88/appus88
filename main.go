package main
import(
"../backend/app"

)

func main() {
	app := app.Startapplication()
}


/*
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":30000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
*/