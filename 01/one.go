package main

import "fmt"
import "io"

func main() {
	prev := 1000000000
	ans := 0
	for true {
		var num int
		_, err := fmt.Scanf("%d", &num)
		if err == io.EOF {
			break
		}
		if num > prev {
			ans += 1
		}
		prev = num
	}
	fmt.Printf("%d\n", ans)
}
