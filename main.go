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



	//slices in golang
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)
	


// my name is ritankar 







	fmt.Println("End of the program.")

	fmt.Println("This is a simple Go program demonstrating basic syntax, loops, and maps.")

	

}

