package message

/**
 * 消息产生器
 */
type MessageProducer interface {

	/**
	 * 添加消息监听器
	 */
	AddMessageListener(listener *MessageListener)

	/**
	 * 移除消息监听器
	 */
	RemoveMessageListener(listener *MessageListener)

	/**
	 * 广播发送消息给消息监听器
	 */
	SendMessage(message *Message)
}
