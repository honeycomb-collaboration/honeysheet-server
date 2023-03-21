package route

import (
	"github.com/gin-gonic/gin"
	"honeysheet-server/web/controller"
	"honeysheet-server/web/websocket"
)

func RegisterRoute() *gin.Engine {
	r := gin.Default()
	group := r.Group("api/v1")

	// get a spreadsheet
	group.GET("/spreadsheet/:spreadsheetId", controller.QuerySpreadSheet)
	// get a spreadsheet's sheets
	group.GET("/spreadsheet/:spreadsheetId/sheet", controller.QuerySpreadSheet)
	// get a sheet
	group.GET("/sheet/:sheetId", controller.QuerySpreadSheet)
	// get a sheet's columns
	group.GET("/sheet/:sheetId/column", controller.QuerySpreadSheet)
	// get a sheet's specific column
	group.GET("/sheet/:sheetId/column/:columnId", controller.QuerySpreadSheet)
	// get a sheet's rows
	group.GET("/sheet/:sheetId/row", controller.QuerySpreadSheet)
	// get a sheet's specific row
	group.GET("/sheet/:sheetId/row/:rowId", controller.QuerySpreadSheet)
	// get a sheet's cells
	group.GET("/sheet/:sheetId/cell", controller.QuerySpreadSheet)
	// get a sheet's specific cell
	group.GET("/sheet/:sheetId/cell/:cellId", controller.QuerySpreadSheet)

	// create one spreadsheet
	group.POST("/spreadsheet", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	group.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	group.Any("/ws", func(c *gin.Context) {
		websocket.Run(c.Writer, c.Request)
	})

	return r
}
