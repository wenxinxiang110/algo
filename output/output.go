package output

import (
	"fmt"
	"math"
)

func OutputDemo() {
	// 十进制整数 %d
	// 十六进制 %x
	// 八进制 %0
	// 二进制 %b
	fmt.Printf("%d\t%x\t%o\t%b\n", 31, 31, 31, 31)

	//	浮点数 %f %g %e
	fmt.Printf("%f\t%g\t%e\n", math.Pi, math.Pi, math.Pi)

	//	 布尔 %t
	fmt.Printf("%t\t%t\n", true, false)

	//	%c          字符（rune） (Unicode码点)
	s := "string"
	fmt.Printf("%c\n", s[0])

	//%s          字符串
	fmt.Printf("%s\n", s)

	//%q          带双引号的字符串"abc"或带单引号的字符'c'
	fmt.Printf("%q\t%q\n", s, s[1])

	//%v          变量的自然形式（natural format）啥子都能输出,但是s[0]这种会输出unicode码..
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t\n", 31, math.Pi, true, s[0], s)

	//%T          变量的类型
	fmt.Printf("%T\t%T\t%T\t%T\t%T\n", 31, math.Pi, true, s[0], s)

	//%%          字面上的百分号标志（无操作数）

}
