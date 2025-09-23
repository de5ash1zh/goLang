package main

import "fmt"

func main (){
	// var name string = "golang"
	// fmt.Println(name);

	var isAdult bool = true
	fmt.Println(isAdult);

	var age int = 30
	fmt.Println(age);

	//shorthand syntax
	name:="golang"
	fmt.Println(name);

	var newName string
	newName = "python"
	fmt.Println(newName);

	var height float64 = 5.9
	fmt.Println(height);

	var complexNum complex128 = 1 + 2i
	fmt.Println(complexNum);

	var char rune = 'A'
	fmt.Println(char);

	var byteValue byte = 65
	fmt.Println(byteValue);
}