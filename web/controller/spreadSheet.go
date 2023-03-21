package controller

import (
	"github.com/gin-gonic/gin"
	"honeysheet-server/internal/service/spreadsheet"
	"net/http"
)

func QuerySpreadSheet(c *gin.Context) {
	ID := c.Param("spreadsheetId")
	ss := spreadsheet.SpreadSheet{}
	err := ss.GetByID(ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, ss)
}
