package main

import (
	"github.com/NothingXiang/algo/img"
	"github.com/NothingXiang/algo/output"
	"math/rand"
	"os"
	"time"
)

func main() {
	output.OutputDemo()
	/*
		// 读取输入的demo
		counts := make(map[string]int)

		input := bufio.NewScanner(os.Stdin)

		//todo: ctrl+d结束输入（其他方式不太清楚能否生效，例如EOF,ctrl+z等）
		for input.Scan() {
			counts[input.Text()]++
		}

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	*/

	rand.Seed(time.Now().UTC().UnixNano())
	file, _ := os.Create("gopl.gif")

	img.Lissajous(file)

	file.Close()
}
