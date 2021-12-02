package main

import "fmt"
import "io"

func main() {
	position:= 0
	depth := 0
	for true {
		var x int
		var direction string
		_, err := fmt.Scanf("%s %d", &direction, &x)
		if err == io.EOF {
			break
		}
		if direction == "forward" {
			position += x
		} else if direction == "up" {
			depth -= x
		} else {
			depth += x
		}
	}
	fmt.Printf("%d\n", position * depth)
}
