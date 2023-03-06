package lexer

import "strings"

type StringSource struct {
	sourceStringList []string // 拆分后的源String
	line             string   // 当前行内容
	lineNum          int      // 当前行号
	currentPos       int      // 当前位置
	printLine        int      // 表示当前已经打印的行号，如果还没有打印为-1
	sourcStr         string
}

func NewStringSource(sourcStr string) *StringSource {
	return &StringSource{
		lineNum:          0,
		currentPos:       -2,
		printLine:        -1,
		sourceStringList: strings.Split(sourcStr, string(EOL)),
		sourcStr:         sourcStr,
	}
}
