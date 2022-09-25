package model

type SpreadSheet struct {
	ID         string `bson:"_id,omitempty"`
	Name       string `form:"name" bson:",omitempty" json:",omitempty"`
	CreateTime string `form:"create_time" bson:",omitempty" json:",omitempty"`
	DeleteTime string `form:"delete_time" bson:",omitempty" json:",omitempty"`
	Deleted    int8   `json:"deleted"`
	Config     interface{}
	Columns    map[string]Column `form:"columns" bson:",omitempty" json:",omitempty"`
}
