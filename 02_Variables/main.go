package main

import "fmt"

func main() {
	var username string = "Sai Mohan"
	fmt.Println(username)
	fmt.Printf("Type of username :  %T \n ", username)

	var rollNo int = 25
	fmt.Println(rollNo)
	fmt.Printf("Type of rollNo :  %T \n ", rollNo)

	var pwAuth bool = true
	fmt.Println(pwAuth)
	fmt.Printf("Type of pwAuth :  %T \n ", pwAuth)

	//implicit type

	var first = 37
	fmt.Println(first)
	fmt.Printf("Type of first : %T \n", first)

	var second = "Sai"
	fmt.Println(second)
	fmt.Printf("Type of second : %T \n", second)

	//default values

	var abc string
	fmt.Println(abc)

	var def int
	fmt.Println(def)

	var ghi bool
	fmt.Println(ghi)

	//no var style

}
