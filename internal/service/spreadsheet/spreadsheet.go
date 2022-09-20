package spreadsheet

import (
	"spreadsheet-server/internal/service/sheet"
)

type SpreadSheet struct {
	ID      string
	Name    string
	Sheets  sheet.Sheet
	Configs interface{}
}

func (*SpreadSheet) Get(ID string) (SpreadSheet, error) {
	ss := SpreadSheet{}
	// todo query spreadsheet
	return ss, nil
}
