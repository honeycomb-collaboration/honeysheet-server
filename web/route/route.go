package route

import (
	"github.com/gin-gonic/gin"
	"spreadsheet-server/web/websocket"
)

func RegisterRoute() *gin.Engine {
	r := gin.Default()
	gv1 := r.Group("v1")
	gv1.Use()

	gv1.GET("/load", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	gv1.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	gv1.Any("/ws", func(c *gin.Context) {
		websocket.Run(c.Writer, c.Request)
	})

	return r
}
