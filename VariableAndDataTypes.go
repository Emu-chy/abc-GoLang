package main

import "fmt"

func VariableAndDataTypeDefinition() {
	var firstName string = "Emon"
	var lastName string = "Chowdhury"
	age := 26
	height := 5.7

	fmt.Println(firstName + lastName)
	fmt.Println(age)
	fmt.Println(height)

	// Data types

	fmt.Printf("the data type of Name is %T\n", firstName+lastName)
	fmt.Printf("the data type of age is %T\n", age)
	fmt.Printf("the data type of height is %T\n", height)

}


/*
## Component           Description
**Variable**           Variable is a box or container, which is contain value(name, text and number) inside. 

**DataTypes**          A data type tells the computer what kind of value you're storing in a variable — like text, number, true/false, etc.


**🛠️ Summary of Variable**
**Use var** when you want to be clear about the type.
**Use :=** when you want Go to figure it out.
Variables store data like text, numbers, or true/false.
Use **fmt.Println()** to show variable values.


**🧠 Summary of DataTypes**
Data types define what kind of value a variable can hold.
Common types: **string, int, float64, bool**
Go is strongly typed, so type matters!
You can use var or := when declaring variables.
*/