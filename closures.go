def'n:
	- closure property is used extensively where data isolation is required. The state provided by the closures makes them immensely helpful in that regard. When we want to create a state encapsulated function we use closures to do that.

	- closure can be thought of as method that takes a method as a parameter.  the method parameter 
	-	a closure is a record storing a function together with an environment

	- whereas an object is data that has one or more functions bound to it, a closure is a function that has one or more variables bound to it.

	- closure ias a function inside of another function.  in java, the innter function would be defined as an interface, then at runtime, the concrete impl would be passed in.
		- The closure can refer to any of the outer function's local variables
		-	the compiler detects that and moves these variables from the outer function's stack space to the closure's hidden object declaration. 
		- You then have a variable of a closure type
			even though it's basically an object under the hood, you pass it around as a function reference, because the focus is on its nature as a function.	

	-  You can see a closure as an object with only one method, and an arbitrary object as a collection of closures over some common underlying data (the object's member variables)

~~~~~ CODE EXAMPLE 

// feels like a java interface
type mathoperation func(int, int) int

func add() mathoperation {
	c := 10
	return func(a int, b int) int {
		return a + b + c
	}
}

func sub() mathoperation {
	c := 10
	return func(a int, b int) int {
		return c - a - b
	}
}

func add_with_parameter(d int) mathoperation {
	c := 10
	return func(a int, b int) int {
		return a + b + c + d
	}
}

func sub_with_parameter(d int) mathoperation {
	c := 10
	return func(a int, b int) int {
		return d - c - a - b
	}
}

func main() {
	adder := add() // or var adder add()
	fmt.Println(adder(2, 2))
	subtractor := sub() // or var subtractor sub()
	fmt.Println(subtractor(2, 2))

	fmt.Println(add_with_parameter(100)(2, 2))
	fmt.Println(sub_with_parameter(100)(2, 2))
}
