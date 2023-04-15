package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"honeysheet-server/web/controller"
	"honeysheet-server/web/websocket"
)

func RegisterRoute() *gin.Engine {
	r := gin.Default()
	group := r.Group("api/v1")

	group.Use(cors.Default())

	// get a workbook
	group.GET("/workbook/:workbookId", controller.GetWorkbook)
	// get a workbook's sheets
	group.GET("/workbook/:workbookId/sheet", controller.GetWorkbookSheets)
	// get a sheet
	group.GET("/sheet/:sheetId", controller.GetWorkbook)
	// get a sheet's columns
	group.GET("/sheet/:sheetId/column", controller.GetWorkbook)
	// get a sheet's specific column
	group.GET("/sheet/:sheetId/column/:columnId", controller.GetWorkbook)
	// get a sheet's rows
	group.GET("/sheet/:sheetId/row", controller.GetWorkbook)
	// get a sheet's specific row
	group.GET("/sheet/:sheetId/row/:rowId", controller.GetWorkbook)
	// get a sheet's cells
	group.GET("/sheet/:sheetId/cell", controller.GetWorkbook)
	// get a sheet's specific cell
	group.GET("/sheet/:sheetId/cell/:cellId", controller.GetWorkbook)
	// create one workbook
	group.POST("/workbook", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	group.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	group.GET("/ws", func(c *gin.Context) {
		websocket.Run(c.Writer, c.Request)
	})

	return r
}
