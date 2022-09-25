package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spreadsheet-server/internal/service/spreadsheet"
)

func QuerySpreadSheet(c *gin.Context) error {
	ID := c.Param("spreadsheetId")
	ss := spreadsheet.SpreadSheet{}
	err := ss.GetByID(ID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, ss)
	return nil
}
