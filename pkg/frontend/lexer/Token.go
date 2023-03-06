package lexer

type Token struct {
	Text     string    // 字面文本
	Value    any       // 值，如果是一些常量，直接可以算出值来的
	Source   Source    // 源文件
	lineNum  int       // 所在行
	position int       // Token第一个字符所在的位置，即行中列位置
	endPos   int       // Token最后一个字符所在的位置，即行中列位置
	Type     TokenType // 语言相关的Token类型
	Note     string    //注释信息
}
