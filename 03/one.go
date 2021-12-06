package main

import "fmt"
import "io"

func main() {
	var s []string
	for true {
		var bits string
		_, err := fmt.Scanf("%s", &bits)
		if err == io.EOF {
			break
		}
		s = append(s, bits)
	}

	freq0 := make([]int, len(s[0]))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == '0' {
				freq0[j]++;
			}
		}
	}

	gamma := 0
	epsilon := 0
	for j := 0; j < len(freq0); j++ {
		if freq0[j] > len(s) - freq0[j] {
			epsilon += (1 << (len(s[0]) - j - 1))
		} else {
			gamma += (1 << (len(s[0]) - j - 1))
		}
	}

	fmt.Printf("%d\n", gamma * epsilon)
}
