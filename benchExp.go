package goBenchExp

import "fmt"

func inlineFunc(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func emptyFunc(i int) {
	if i == 555 {
		fmt.Println(i)
	}
}

func simpleDefer(i int) (j int) {
	defer func() {
		j = i + 1
	}()
	return
}
