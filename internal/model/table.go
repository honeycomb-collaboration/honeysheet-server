package model

import (
	"spreadsheet-server/internal/model/column"
)

type Table struct {
	Id         string `json:"_id"`
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
	DeleteTime string `json:"delete_time"`
	Deleted    string `json:"deleted"`
	Columns    struct Columns
	Config struct {
	} `json:"config"`
}