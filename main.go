package main

import "fmt"


func main(){


	fmt.Println("Hello, World!")


	//looping in golang 
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}


	// maps in golang
	person := make(map[string]string)
	person["name"] = "John"
	person["age"] = "30"
	person["city"] = "New York"

	fmt.Println("Person:", person)





	fmt.Println("End of the program.")

	

}