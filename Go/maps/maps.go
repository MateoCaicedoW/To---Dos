package main

import "fmt"

func main() {
	dicc := map[string]string{
		"go": "ir",
	}

	if dicc["go"] == "ir" {
		fmt.Println("go es ir")
		return
	}
	fmt.Println("no existe")

}
