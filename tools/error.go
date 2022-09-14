package tools
// 定义原始错误

type ErrorMsg struct {
	ErrorCode int    `json:"ErrorCode"`
	Message   string `json:"Message"`
}

func (e ErrorMsg) Error() string {
	return e.Message
}

var (
	Error                   ErrorMsg
	ErrorXXX           		 ErrorMsg
)

func init() {

	Error = ErrorMsg{ErrorCode: 30000, Message: "通用错误"}
	ErrorXXX = ErrorMsg{ErrorCode: 30001, Message: "例子"}

}
