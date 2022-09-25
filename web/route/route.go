package route

import (
	"github.com/gin-gonic/gin"
	"spreadsheet-server/web/controller"
	"spreadsheet-server/web/websocket"
)

func RegisterRoute() *gin.Engine {
	r := gin.Default()
	gv1 := r.Group("v1/spreadsheets")
	gv1.Use()

	// get one spreadsheet's all data
	gv1.GET("/get/:spreadsheetId", func(c *gin.Context) {
		err := controller.QuerySpreadSheet(c)
		defer func() {
			if err != nil {
				panic(err)
			}
		}()
	})

	// create one spreadsheet
	gv1.POST("/create", func(c *gin.Context) {
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
