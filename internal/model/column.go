package model

type Column struct {
	V          string `json:"v"`
	O          string `json:"o"`
	Deleted    string `json:"deleted"`
	CreateTime string `json:"create_time"`
	DeleteTime string `json:"delete_time"`
}
