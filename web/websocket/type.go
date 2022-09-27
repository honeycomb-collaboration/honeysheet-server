package websocket

type MessageType string // basic spreadsheet command

const (
	OpenMessage      MessageType = "OPEN_MESSAGE"      // open one spreadsheet
	CloseMessage     MessageType = "CLOSE_MESSAGE"     // close one
	OperationMessage MessageType = "OPERATION_MESSAGE" // other operation
	CursorMessage    MessageType = "CURSOR_MESSAGE"    // cursor operation
)

type OperationType string // Some spreadsheet operations that will change spreadsheet data.

const (
	AddField      OperationType = "AddField"      // add columns
	AddRecords    OperationType = "AddRecords"    // add rows
	DeleteField                 = "DeleteField"   // delete columns
	DeleteRecords               = "DeleteRecords" // delete rows
	SetFieldAttr                = "SetFieldAttr"  // update column
	SetRecord                   = "SetRecord"     // update rows
	MoveRecords                 = "MoveRecords"   // move rows
	// todo other operations
)

type Action struct {
	TID  string      // spreadsheet ID
	CID  string      // column ID
	RID  string      // row ID
	data interface{} // operation special data
}

type OperationStruct struct {
	Command OperationType
	Action  Action
}

// MessageStruct websocket request body
type MessageStruct struct {
	RequestID          string
	Type               MessageType
	Operations         []OperationStruct
	Token              string
	__InjectParameters struct { // 在逐层过滤中，给message注入结构，方便后面逻辑使用
		User         interface{}
		HeaderParams interface{}
	}
}
