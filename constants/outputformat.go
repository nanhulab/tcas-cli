/*
 * @Author: fanjf
 * @Date: 2024-07-26 11:01:16
 * @LastEditTime: 2024-07-26 11:17:47
 * @LastEditors: fanjf
 * @FilePath: \tcas-cli\constants\outputformat.go
 * @Description: 🎉🎉🎉
 */
package constants

const (
	OutReset = "\033[0m"
	//颜色
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
	//字体样式
	FontBold         = "\033[1m"
	FontUnderline    = "\033[4m"
	FontItalic       = "\033[3m"
	FontBlink        = "\033[5m"
	FontBoldOff      = "\033[21m"
	FontUnderlineOff = "\033[24m"
	FontItalicOff    = "\033[23m"
	FontBlinkOff     = "\033[25m"
	//背景颜色
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgPurple  = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
	BgDefault = "\033[49m"
	//光标
	CursorOff            = "\033[?25l"
	CursorOn             = "\033[?25h"
	CursorUp             = "\033[A"
	CursorDown           = "\033[B"
	CursorRight          = "\033[C"
	CursorLeft           = "\033[D"
	CursorHome           = "\033[H"
	CursorEnd            = "\033[F"
	CursorSave           = "\033[s"
	CursorRestore        = "\033[u"
	CursorClear          = "\033[J"
	CursorClearDown      = "\033[J"
	CursorClearUp        = "\033[1J"
	CursorClearAll       = "\033[2J"
	CursorClearLine      = "\033[2K"
	CursorClearLineLeft  = "\033[1K"
	CursorClearLineRight = "\033[K"
	CursorClearLineAll   = "\033[3K"
)
