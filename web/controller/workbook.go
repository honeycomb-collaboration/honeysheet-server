package controller

import (
	"github.com/gin-gonic/gin"
	"honeysheet-server/internal/service/workbook"
	"net/http"
)

func GetWorkbook(c *gin.Context) {
	ID := c.Param("workbookId")
	w := workbook.Workbook{}
	err := w.GetByID(ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, w)
}
