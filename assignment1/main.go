package main

import "log"

func main() {
	nums := make([]int, 10)

	for n := range nums {
		if n%2 == 0 {
			log.Println(n, " is odd")
		} else {
			log.Println(n, " is even")
		}
	}
}
