package main

import "fmt"
import "io"

func main() {
	var numbers []int
	for true {
		var num int
		_, err := fmt.Scanf("%d", &num)
		if err == io.EOF {
			break
		}
		numbers = append(numbers, num)
	}

	length := len(numbers)
	ans := 0
	for i := 1; i + 2 < length; i++ {
		prev := numbers[i - 1] + numbers[i] + numbers[i + 1]
		cur := numbers[i] + numbers[i + 1] + numbers[i + 2]
		if cur > prev {
			ans += 1
		}
	}
	fmt.Printf("%d\n", ans)
}
