package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//n := 4

	list := [7]string{"negro", "rojo", "azul", "verde", "amarillo", "rosado", "violeta"}
	var num int
	for i := 0; i < 10; i++ {
		num = rand.Intn(len(list))

	}
	Color(list, num)

	//

	/*
		for true {
			n++
			if n >= 12 {
				break
			}
		}

		fmt.Println(n)
	*/

}

func Color(list [7]string, num int) {
	if num%2 == 0 {
		fmt.Println(list[num])
	} else {
		Color(list, num)
	}

}
