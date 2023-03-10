package message

// 消息类型
type MessageType int

const (
	SOURCE_LINE MessageType = iota
	SYNTAX_ERROR
	PARSER_SUMMARY
	INTERPRETER_SUMMARY
	COMPILER_SUMMARY
	MISCELLANEOUS
	TOKEN
	ASSIGN
	FETCH
	BREAKPOINT
	RUNTIME_ERROR
	CALL
	RETURN
)
