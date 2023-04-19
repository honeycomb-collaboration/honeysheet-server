package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Workbook struct {
	ID                 string `json:"id"`
	Name               string `json:"name,omitempty"`
	DefaultColumnCount uint16 `json:"defaultColumnCount,omitempty"` // 默认列数
	DefaultRowCount    uint16 `json:"defaultRowCount,omitempty"`    // 默认行数
	DefaultColumnWidth uint16 `json:"defaultColumnWidth,omitempty"` // 默认列宽
	DefaultRowHeight   uint16 `json:"defaultRowHeight,omitempty"`   // 默认行高
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
	ColumnIds          []ColumnId `json:"columnIds"`
	RowIds             []RowId    `json:"rowIds"`
	Name               string     `json:"name,omitempty"`
	DefaultColumnWidth uint16     `json:"defaultColumnWidth,omitempty"` // 默认列宽
	DefaultRowHeight   uint16     `json:"defaultRowHeight,omitempty"`   // 默认行高
}

var sheet = Sheet{
	ID:        "1",
	Name:      "Sheet 1",
	ColumnIds: []ColumnId{0, 1, 2, 3, 4},
	RowIds:    []RowId{0, 1, 2, 3, 4},
}

var sheets = []*Sheet{
	&sheet,
}

type Cell struct {
	V string `json:"v"`
}

type CellRecord struct {
	ColumnId ColumnId `json:"columnId"`
	RowId    RowId    `json:"rowId"`
	Cell     *Cell    `json:"cell"`
}

var cell = Cell{
	V: "R0_C0",
}

var cells = []CellRecord{
	{
		ColumnId: 0,
		RowId:    0,
		Cell:     &cell,
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

func GetSheet(c *gin.Context) {
	ID := c.Param("sheetId")
	fmt.Println("get sheet by id", ID)

	c.JSON(http.StatusOK, sheet)
}

func GetSheetCells(c *gin.Context) {
	ID := c.Param("sheetId")
	fmt.Println("get sheet's cells by id", ID)

	c.JSON(http.StatusOK, cells)
}
func GetCell(c *gin.Context) {
	sheetId := c.Param("sheetId")
	cellId := c.Param("cellId")
	fmt.Println("get cell by id", sheetId, cellId)

	c.JSON(http.StatusOK, cell)
}

func UpdateCell(v string) {
	fmt.Println("set cell with", v)
	cell.V = v
}

func TODO(c *gin.Context) {
	c.JSON(http.StatusNotFound, nil)
}
