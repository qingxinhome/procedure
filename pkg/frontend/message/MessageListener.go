package message

/**
 * 消息监听器
 */
type MessageListener interface {

	/**
	 * 接收并处理消息
	 */
	MessageReceived(message Message)
}
