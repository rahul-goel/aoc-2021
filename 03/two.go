package main

import "fmt"
import "io"

func remove(s []int, i int) []int {
	s[i] = s[len(s) - 1]
	return s[:len(s) - 1]
}

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

	t := make([]string, len(s))
	copy(t, s)

	for bit := 0; bit < len(s[0]); bit++ {
		f0 := 0
		f1 := 0
		for i:= 0; i < len(s); i++ {
			if s[i][bit] == '0' {
				f0++;
			} else {
				f1++;
			}
		}
		rmv := 1
		if (f0 > f1) {
			rmv = 0
		}
		for i := 0; i < len(s); i++ {
			if int(s[i][bit]) - int('0') == rmv {
				s = append(s[:i], s[i + 1:]...)
				i--
			}
		}
		if len(s) == 1 {
			break;
		}
	}

	for bit := 0; bit < len(t[0]); bit++ {
		f0 := 0
		f1 := 0
		for i:= 0; i < len(t); i++ {
			if t[i][bit] == '0' {
				f0++;
			} else {
				f1++;
			}
		}
		rmv := 0
		if (f1 < f0) {
			rmv = 1
		}
		for i := 0; i < len(t); i++ {
			if int(t[i][bit]) - int('0') == rmv {
				t = append(t[:i], t[i + 1:]...)
				i--
			}
		}
		if len(t) == 1 {
			break;
		}
	}
	num0 := 0
	for i := 0; i < len(s[0]); i++ {
		num0 += (1 << (len(s[0]) - i - 1) * (int(s[0][i]) - int('0')))
	}
	num1 := 0
	for i := 0; i < len(t[0]); i++ {
		num1 += (1 << (len(t[0]) - i - 1) * (int(t[0][i]) - int('0')))
	}

	fmt.Printf("%d\n", num0 * num1)

}
