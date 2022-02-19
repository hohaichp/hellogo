package color

import (
	"fmt"
	"strconv"
)

type Color uint8

// Foreground colors. basic foreground colors 30 - 37
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta // 品红
	FgCyan    // 青色
	FgWhite
	// FgDefault revert default FG
	FgDefault Color = 39
)

// Extra foreground color 90 - 97(非标准)
const (
	FgDarkGray Color = iota + 90 // 亮黑（灰）
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgLightWhite
	// FgGray is alias of FgDarkGray
	FgGray Color = 90 // 亮黑（灰）
)

// Background colors. basic background colors 40 - 47
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow // BgBrown like yellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// BgDefault revert default BG
	BgDefault Color = 49
)

// Extra background color 100 - 107(非标准)
const (
	BgDarkGray Color = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgLightWhite
	// BgGray is alias of BgDarkGray
	BgGray Color = 100
)

// Option settings
const (
	OpReset         Color = iota // 0 重置所有设置
	OpBold                       // 1 加粗
	OpFuzzy                      // 2 模糊(不是所有的终端仿真器都支持)
	OpItalic                     // 3 斜体(不是所有的终端仿真器都支持)
	OpUnderscore                 // 4 下划线
	OpBlink                      // 5 闪烁
	OpFastBlink                  // 5 快速闪烁(未广泛支持)
	OpReverse                    // 7 颠倒的 交换背景色与前景色
	OpConcealed                  // 8 隐匿的
	OpStrikethrough              // 9 删除的，删除线(未广泛支持)
)

// There are basic and light foreground color aliases
const (
	Red     = FgRed
	Cyan    = FgCyan
	Gray    = FgDarkGray // is light Black
	Blue    = FgBlue
	Black   = FgBlack
	Green   = FgGreen
	White   = FgWhite
	Yellow  = FgYellow
	Magenta = FgMagenta

	// special

	Bold   = OpBold
	Normal = FgDefault

	// extra light

	LightRed     = FgLightRed
	LightCyan    = FgLightCyan
	LightBlue    = FgLightBlue
	LightGreen   = FgLightGreen
	LightWhite   = FgLightWhite
	LightYellow  = FgLightYellow
	LightMagenta = FgLightMagenta
)

const (
	SettingTpl   = "\x1b[%sm"
	FullColorTpl = "\x1b[%sm%s\x1b[0m"
)

// String convert to code string. eg "35"
func (c Color) String() string {
	return strconv.Itoa(int(c))
}

// Usage: code := "3;32;45"
func (c Color) Sprintf(format string, args ...interface{}) string {

	code := c.String()
	str := fmt.Sprintf(format, args...)

	if len(code) == 0 || str == "" {
		return str
	}
	return fmt.Sprintf(FullColorTpl, code, str)
}

// 其中0x1B是标记，[开始定义颜色，1代表高亮，40代表黑色背景，32代表绿色前景，0代表恢复默认颜色。显示效果为：
// fmt.Printf("\x1b[1;40;32m%s\x1b[0m", "testPrintColor\r\n")
// code := "3;32;45"
func (c Color) Printf(format string, args ...interface{}) {

	code := c.String()
	str := fmt.Sprintf(format, args...)

	if len(code) == 0 || str == "" {
		fmt.Printf(str)
	}

	fmt.Printf(FullColorTpl, code, str)
}


/**

tips

// 其中0x1B是标记，[开始定义颜色，1代表高亮，40代表黑色背景，32代表绿色前景，0代表恢复默认颜色。显示效果为：
fmt.Printf(color.FullColorTpl, strconv.Itoa(int(color.Normal)), "teststring\r\n");
fmt.Printf(color.FullColorTpl, strconv.Itoa(int(color.Bold)), "teststring\r\n");
fmt.Printf(color.FullColorTpl, strconv.Itoa(int(color.OpItalic)), "teststring\r\n");
fmt.Printf(color.FullColorTpl, strconv.Itoa(90), "teststring\r\n");
fmt.Printf(color.FullColorTpl, strconv.Itoa(91), "teststring\r\n");
fmt.Printf(color.FullColorTpl, strconv.Itoa(int(color.FgGreen)), "teststring\r\n");
for f := 30; f <= 37; f++ {
	fmt.Printf("\x1b[%sm%s\x1b[0m", strconv.Itoa(f), "teststring\r\n");
}

// 其中0x1B是标记，[开始定义颜色，1代表高亮，40代表黑色背景，32代表绿色前景，0代表恢复默认颜色。显示效果为：
fmt.Printf("\x1b[1;40;32m%s\x1b[0m", "testPrintColor\r\n")

temp := []byte(fmt.Sprintf(color, "32", "teststring\r\n"))
os.Stdout.Write(temp)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
	for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
		for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
			fmt.Printf("%c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

*/
