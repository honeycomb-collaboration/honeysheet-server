package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Workbook struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	DefaultColumnCount uint16 `json:"defaultColumnCount"` // 默认列数
	DefaultRowCount    uint16 `json:"defaultRowCount"`    // 默认行数
	DefaultColumnWidth uint16 `json:"defaultColumnWidth"` // 默认列宽
	DefaultRowHeight   uint16 `json:"defaultRowHeight"`   // 默认行高
}

var w = Workbook{
	ID:                 "honeysheet_id",
	Name:               "My Honeysheet",
	DefaultColumnCount: 36,  // 默认列数
	DefaultRowCount:    36,  // 默认行数
	DefaultColumnWidth: 120, // 默认列宽
	DefaultRowHeight:   40,  // 默认行高
}

type ColumnId uint32
type RowId uint32

type Sheet struct {
	ID                 string     `json:"id"`
	Name               string     `json:"name"`
	ColumnIds          []ColumnId `json:"columnIds"`
	RowIds             []RowId    `json:"rowIds"`
	DefaultColumnWidth uint16     `json:"defaultColumnWidth,omitempty"` // 默认列宽
	DefaultRowHeight   uint16     `json:"defaultRowHeight,omitempty"`   // 默认行高
}

var sheets = []Sheet{
	Sheet{
		ID:        "1",
		Name:      "Sheet 1",
		ColumnIds: []ColumnId{0, 1, 2, 3, 4},
		RowIds:    []RowId{0, 1, 2, 3, 4},
	},
}

func GetWorkbook(c *gin.Context) {
	ID := c.Param("workbookId")
	fmt.Println("get workbook by id", ID)
	//w := workbook.Workbook{}
	//err := w.GetByID(ID)
	//if err != nil {
	//	_ = c.AbortWithError(http.StatusInternalServerError, err)
	//}

	c.JSON(http.StatusOK, w)
}

func GetWorkbookSheets(c *gin.Context) {
	ID := c.Param("workbookId")
	fmt.Println("get workbook's sheets by id", ID)

	c.JSON(http.StatusOK, sheets)
}
