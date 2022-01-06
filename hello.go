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
   array()
   interfaces()
}

func interfaces(){
	c1 := car{"suzuki","blue"}  
	t1:= toyota{"Toyota","Red", 100}  
	c1.accelerate()  
	t1.accelerate()  
	foo(c1)  
	foo(t1)  
 
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

func array(){
	var x [5]int
	var i,j int

	for i=0;i<5;i++{
		x[i]=i+10
	}
	for j=0;j<5;j++{
		fmt.Printf("Element[%d]=%d\n",j,x[j])
	}
}

type vehicle interface {  
   accelerate()  
}  
func foo(v vehicle)  {  
   fmt.Println(v)  
     
}  
type car struct {  
   model string  
   color string  
}  
func (c car) accelerate()  {  
   fmt.Println("Accelrating?")  
     
}  
type toyota struct {  
  model string  
   color string  
   speed int  
}  
func (t toyota) accelerate(){  
   fmt.Println("I am toyota, I accelerate fast?")  
}