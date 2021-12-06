package main

import "fmt"
import "io"

func main() {

	var cnt [10] int
	for true {
		var value int
		_, err := fmt.Scanf("%d,", &value)
		if err == io.EOF {
			break
		}
		cnt[value]++
	}

	for i := 0; i < 80; i++ {
		var new_cnt [10] int
		for j := 0; j < 9; j++ {
			if j == 0 {
				new_cnt[8] += cnt[j]
				new_cnt[6] += cnt[j]
			} else {
				new_cnt[j - 1] += cnt[j]
			}
		}
		cnt = new_cnt
	}

	var ans int
	for i := 0; i < 9; i++ {
		ans += cnt[i]
	}

	fmt.Println(ans)
}
