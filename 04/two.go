package main

import "fmt"
import "strings"
import "strconv"
import "io"

func check(vis [5][5] bool, mat [5][5] int) (bool, int) {
	unmarked := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !vis[i][j] {
				unmarked += mat[i][j]
			}
		}
	}
	for i := 0; i < 5; i++ {
		tmp := true
		for j := 0; j < 5; j++ {
			tmp = tmp && vis[i][j]
		}
		if tmp {
			return true, unmarked
		}
	}
	for i := 0; i < 5; i++ {
		tmp := true
		for j := 0; j < 5; j++ {
			tmp = tmp && vis[j][i]
		}
		if tmp {
			return true, unmarked
		}
	}
	return false, 0
}

func main() {
	var s string
	fmt.Scanf("%s\n\n", &s)
	s_slice := strings.Split(s, ",")
	for i := 0; i < len(s_slice); i++ {
		s_slice[i] = strings.TrimSpace(s_slice[i])
	}

	var nums []int
	for i := 0; i < len(s_slice); i++ {
		num, err := strconv.Atoi(s_slice[i])
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	score := 0
	moves := -1

	for true {
		var mat [5][5] int
		var vis [5][5] bool
		done := false
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				_, err := fmt.Scanf("%d", &mat[i][j])
				if err == io.EOF {
					done = true
				}
				if done {
					break
				}
			}
			if done {
				break
			}
		}
		if done {
			break
		}
		fmt.Scanf("\n");
		for k := 0; k < len(nums); k++ {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if mat[i][j] == nums[k] {
						vis[i][j] = true
					}
				}
			}
			res, unmarked := check(vis, mat)
			if res && k + 1 > moves {
				moves = k + 1
				score = unmarked * nums[k]
				break
			} else if res {
				break
			}
		}
		if done {
			break
		}
	}

	fmt.Println(score)
}
