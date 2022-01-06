package main

import "rsc.io/quote"
import "fmt"


type Employee struct {
	fname string
	lname string
}

func (emp Employee)fullname(){
	fmt.Println(emp.fname+""+emp.lname)
}
func main() {
    fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Hello())

	// Data types
	var i int
	var f float64
	var b bool
	var s string

   fmt.Printf("%T %T %T %T\n",i,f,b,s) // Prints type of the variable  
   fmt.Printf("%v   %v      %v  %q     \n", i, f, b, s) //prints initial value of the variable  
   e1:=Employee{"john","Ponting"}
   e1.fullname()
   closure()
}

func closure(){
	number:=10
	squareNumber :=func()(int){
		number*=number
		return number
	}
	fmt.Println(squareNumber())

	fmt.Println(squareNumber())
}