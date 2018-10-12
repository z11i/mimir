package main

import "fmt"

type color string

var (
	bBlack  color = "0;90"
	red     color = "0;31"
	blue    color = "0;34"
	magenta color = "0;35"
	cyan    color = "0;36"
)

func withColor(c color, str string) string {
	return fmt.Sprintf("\x1B[%sm%s\x1B[0m", c, str)
}
