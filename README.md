## Component	    Description
package main	Defines the package as the main entry point of the application

import "fmt"	Imports the fmt package for formatted I/O (e.g., printing)

func main()	Main function where execution starts

fmt.Println()	Prints the string followed by a newline


## package main 
What it means: This tells Go that this file belongs to the main package.
Why it's important: In Go, the main package is special – it tells the compiler that this is the entry point of your program.
In short: Any Go program that you want to run (not a library) must start with package main.

## import "fmt"
What it means: This imports the fmt package, which provides functions for formatted I/O (like printing text).
Common functions: Print(), Println(), Printf().
Here: We're using fmt.Println to print a string.

## func main()
What it means: This defines the main function, the starting point of execution in Go.
Special role: Go always starts running from the main() function in the main package.