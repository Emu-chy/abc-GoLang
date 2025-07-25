## Component	    Description
**package main**	Defines the package as the main entry point of the application

**import "fmt"**	Imports the fmt package for formatted I/O (e.g., printing)

**func main()**	Main function where execution starts

**fmt.Println()**	Prints the string followed by a newline


## package main 
**🧠What it means:** This tells Go that this file belongs to the main package.<br>
**🎯 Why it's important:** In Go, the main package is special – it tells the compiler that this is the entry point of your program.<br>
**🏁In short:** Any Go program that you want to run (not a library) must start with package main.

## import "fmt"
**🧠What it means:** This imports the fmt package, which provides functions for formatted I/O (like printing text).<br>
**Common functions:** Print(), Println(), Printf().<br>
**Here:** We're using fmt.Println to print a string.

## func main()
**🧠What it means:** This defines the main function, the starting point of execution in Go.<br>
**🎯Special role:** Go always starts running from the main() function in the main package.




## Documentation Part 

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


**Conditional Statement**
A conditional statement allows your program to make decisions based on certain conditions.
It helps the program choose what to do next depending on true or false results.

**Syntex of conditional Statement**
if condition1 {
    // runs if condition1 is true
} else if condition2 {
    // runs if condition2 is true
} else {
    // runs if none of the above are true
}
