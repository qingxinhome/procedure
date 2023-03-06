package lexer

const FileSourceType string = "FILE"
const StringSourceType string = "STRING"

const EOL byte = '\n'           // 行结束符
const EOF byte = byte(0)        //文件结束标识
const ENCODING string = "UTF-8" //编码格式

type Source interface {
	/**
	 * 检查编码格式
	 */
	CheckEncode() (bool, error)

	/**
	 * 获取当前字符
	 */
	CurrentChar() (rune, error)

	/**
	 * 获取下一个字符
	 */
	NextChar() (rune, error)

	/**
	 * 查看下一个字符,并保持当前字符不变
	 */
	PeekChar() (rune, error)

	/**
	 * 查看当前字符后第index字符,保持当前字符不变
	 */
	PeekIndexChar(index int) (rune, error)

	/**
	 * 获取当前行内容
	 */
	getLine() (string, error)

	/**
	 * 设置当前行内容
	 */
	SetLine(line string) error

	/**
	 * 获取当前行号
	 */
	GetLineNum() (int, error)

	/**
	 * 设置当前行号
	 */
	SetLineNum(lineNum int) error

	/**
	 * 获取当前位置
	 */
	GetCurrentPos() (int error)

	/**
	 * 设置当前位置
	 */
	SetCurrentPos(current int) error

	/**
	 * 截取SQL片段
	 * @param fromlineNum 开始行号
	 * @param fromPos 开始列号
	 * @param length 起始token的长度
	 * @param endLineNum 结束行号
	 * @param endPos 结束列号
	 * @return
	 */
	CutSubStr(fromlineNum int, fromPos int, length int, endLineNum int, endPos int) string

	/**
	 * 获取源码片段，行号从1开始，位置从0开始，前后包含
	 */
	GetOriginalStr(fromlineNum int, fromPos int, endLineNum int, endPos int)

	/**
	 * 获取源码剩余片段，行号从fromlineNum开始，位置从fromPos + length开始
	 */
	SubCodeString(fromlineNum int, fromPos int, length int)

	/**
	 * 获取源码剩余片段，行号从fromlineNum开始，位置从fromPos开始
	 */
	GetOriginalRemainStr(fromlineNum int, fromPos int)

	/**
	 * 替换源码字符串中从fromlineNum开始，位置从fromPos 开始第一次出现的单词oldWord 替换为newWord
	 * @param fromlineNum
	 * @param fromPos
	 * @param word
	 * @return
	 */
	RepalceWord(fromlineNum int, fromPos int, oldWord string, newWord string) string
	/**
	 * 读取上一个字母,保持位置不变，行号从1开始，位置从0开始
	 */
	LastChar(line int, pos int) (rune, error)

	/**
	 * 替换一截字符串
	 * @param fromlineNum
	 * @param fromPos
	 * @param endLineNum
	 * @param endPos
	 * @param newStr
	 * @return
	 */
	ReplaceSubStr(fromlineNum int, fromPos int, endLineNum int, endPos int, newStr string)

	/**
	 * 返回语句开始token位置
	 * @param token
	 */
	Backto(token Token)
}
