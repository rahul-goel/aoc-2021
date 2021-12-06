package main

import "fmt"
import "io"

func main() {
	var mat[1000][1000] int
	for true {
		var x1, y1, x2, y2 int
		_, err := fmt.Scanf("%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err == io.EOF {
			break
		}
		fmt.Printf("%d %d %d %d\n", x1, y1, x2, y2)

		if x1 == x2 {
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for i := y1; i <= y2; i++ {
				mat[x1][i] += 1
			}
		} else if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				mat[i][y1] += 1
			}
		}
	}

	cnt := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if mat[i][j] >= 2 {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
