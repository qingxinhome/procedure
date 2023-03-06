package message

/**
 * 消息处理器
 */
type MessageHandler struct {
	Message   Message           //消息
	listeners []MessageListener //消息监听器集合
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		listeners: make([]MessageListener, 0),
	}
}

/**
 * 添加消息监听器
 */
func (handdler *MessageHandler) AddMessageListener(listener MessageListener) {
	handdler.listeners = append(handdler.listeners, listener)
}

/**
 * 移除消息监听器
 */
func (handdler *MessageHandler) RemoveMessageListener(listener MessageListener) {
	j := 0
	for _, v := range handdler.listeners {
		if v != listener {
			handdler.listeners[j] = v
			j++
		}
	}
	handdler.listeners = handdler.listeners[:j]
}

/**
 * 广播发送消息给消息监听器的具体实现
 */
func (handdler *MessageHandler) NotifyListeners(message Message) {
	for _, listener := range handdler.listeners {
		listener.MessageReceived(message)
	}
}

/**
 * 广播发送消息给消息监听器
 */
func (handdler *MessageHandler) SendMessage(message Message) {
	handdler.Message = message
	handdler.NotifyListeners(message)
}
