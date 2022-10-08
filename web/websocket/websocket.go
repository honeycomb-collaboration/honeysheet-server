package websocket

import (
	"fmt"
	"github.com/OnlineCollaboration/connection"
	"net/http"
)

// Run todo 开启长链接服务
func Run(responseWriter http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	err := connection.Handle(ctx, responseWriter, request, func(con *connection.Connection) {
		for msg := range con.IncomingMessage {
			fmt.Printf("websocket message %s | %d Bytes\n", string((*msg)[:]), len(*msg))
			sendMsg := []byte("hello, client")
			con.OutgoingMessage <- &sendMsg
		}
		// todo 1. 在 web/websocket/filter-handler 中采用责任链模式对长链接对请求内容做校验。
		// todo 2. 采用策略模式分发请求参数,在 web/websocket/message 中实现消息大类型的区分。
		// todo 3. 采用策略模式分发请求参数,在 web/websocket/operation 中实现具体操作的区分。
		// todo 4. 请求最终调用 interval/service 下的方法完成实际操作。
	})

	if err != nil {
		fmt.Println("websocket handle error", err)
		return
	}
}
