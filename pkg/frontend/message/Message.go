package message

/**
 * 消息
 */
type Message struct {
	//表示当前已打印的行号，如果还没有打印为-1
	PrintLine int
	Body      any // 消息内容
	Level     MessageLevel
	MsgType   MessageType //消息类型
}

func NewMessage(printLine int, body any, msgType MessageType) *Message {
	return &Message{
		PrintLine: printLine,
		Body:      body,
		MsgType:   msgType,
	}
}

type MessageLevel int

const (
	DEBUG MessageLevel = iota
	INFO
	OVERWRITE
	WARN
	ERROR
)
