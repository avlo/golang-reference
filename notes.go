goshimmer - node software allowing nodes to reach consensus without coordinator
wasp - 2nd layer node software atop which run smart contracts
pollen - testnet

~~~

go-lang
https://play.golang.org/

https://github.com/iotaledger/wasp/blob/develop/packages/vm/examples/donatewithfeedback/dwfimpl/impl.go
https://github.com/iotaledger/wasp/blob/master/articles/intro/dwf.md

~~~~
import variations
	- aliased import
	  import f "fmt"
		  "f" can be used in place of fmt, ex:
				f.Println("stuff")
  - dot import
    import . "fmt"
			allows function call without package name, ex:
				Println("stuff")
	- blank import
		import _ "fmt"
			allows for ignoring unused imports
  - nested import
		import "math/rand"
			import subpackage of larger package

~~~~

fmt.Printf("%q\n", i) // quotes for strings, quoted hex for int
fmt.Printf("%s\n", i) // string value
fmt.Printf("%v\n", i) // value
fmt.Printf("%x\n", i) // hex rep

~~~~
arrays and maps are always passed by value
	but only array setting a variable equals copies, while maps internally use pointers
slices are passed by reference
	setting a variable equal to slice and map creates a pointer to  original

~~~~

install all packages in the current directory and its subdirectories
  go install ./...

test all packages in the current directory and its subdirectories
  go test ./...

~~~~

pointer basics

var x = 100
var p *int = &x  // *int is a pointer to an int.  an address of an int is assigned to it
  or shorcut
var p = &x  // infers p is a pointer to an int
*p = 200 // change the integer value pointed to by p  

func zeroptr(iptr *int) {  // *int parameter == parameter takes an int pointer
	*iptr = 0  // *iptr then dereferences the pointer to current value at that address
}

i := 1
zeroptr(&i) // call method with pointer to i.  aka, i's memory address

~~~~

type keyword - https://golang.org/ref/spec#Type_definitions

https://stackoverflow.com/questions/53689968/what-exactly-does-the-type-keyword-do-in-go

type <new_type> <existing_type or type_definition>

Create a type while defining struct.
Format:
	type <new_type> struct { 
		/*...*/
	}

~~~~

function - does not belong to any struct type

	Define function type, (aka. by assigning name to a function signature).
		Format:
			type <FuncName> func(<param_type_list>) <return_type>
		e.g
			type AdderFunc func(int, int) int

methods - methods belong to struct type as per OO

	type post struct {
		title string
	}

	func (p post) details() {  // method works on class of type post
	  fmt.Println("Title: ", p.title)
	}

	func details(p post) {  // function works on parameter of type post
	  fmt.Println("Title: ", p.title)
	}

	func main() {
	  post1 := post{
	    "Inheritance in Go",
	  }
	  post1.details()

	  details(post1)
	}


~~~~

inheritance performed through composition

type author struct {
  firstName string
}

type post struct {
  title     string
  author
}

func (p post) details() {  // method works on class of type post
  fmt.Println("Title: ", p.title)
  fmt.Println("Author: ", p.firstName) // works like inheritance
}

func main() {
  author1 := author{
   "Naveen",
  }

  post1 := post{
    "Inheritance in Go",
    author1,
  }
  post1.details() // method works on class of type post
}

~~~~

pointers as parameters, simplest case
	object is passed by reference(&) and inside the function dereferenced using *

func changeValue(num int, newnum int) {
  num = newnum
}

func changeValueByRef(num *int, newnum int) { // pointer to int
  *num = newnum // pointer dereference
}

func main() {
  val := 12
  fmt.Printf("The value before function call is %d\n", val)
  changeValue(val, 13)         // does not change value
  fmt.Printf("The value after function call is %d\n", val)
  changeValueByRef(&val, 14)   // pass object by reference
  fmt.Printf("The value after function call is %d\n", val)
}

~~

pointers on class reference, next simplest case
	methods on a class use receiver type of pointer

package main

import (
  "fmt"
)

type MyType struct {
  age string
}

func (m *MyType) setMyTypeByRef(input string) {
  m.age = input
}

func (m MyType) setMyTypeByValue(input string) {
  m.age = input
}

func main() {
  myType := MyType{"12"}
  fmt.Println(myType.age) // will be 12
  myType.setMyTypeByValue("13") // try set to 13
  fmt.Println(myType.age) // does not set to 13 as expected
  myType.setMyTypeByRef("13") // or (&myType).setMyType("13")
  fmt.Println(myType.age)
}

~~~~

constructors / new keyword

package main

import (
	"fmt"
)

func main() {
	var m *myStruct
	m = &myStruct{42}
	fmt.Println(m.attr)

	// same as

	var n *myStruct
	n = new(myStruct)
	(*n).attr = 32
	fmt.Println((*n).attr)

	// same as

	var o *myStruct
	o = new(myStruct)
	o.attr = 52
	fmt.Println(o.attr)
}

type myStruct struct {
	attr int
}

~~~~

maps:

	map[<key-type>]value-type
	ex:
		m := map[string]int {
			"ca": 123,
			"tx": 456
		}

	m := map[[]int]string // not allowed, a slice ([]<type>) cannot be a key to a map
	m := map[[3]int]string // is allowed, because array can be a key to a map

	so to populate map dynamically/loop:
		m := make(map[string]int)

package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := map[string]int{
		"ca": 123,
		"tx": 456,
	}
	fmt.Println(m)

	ma := make(map[int]int)
	for i := 0; i <= 9; i++ {
		ma[i] = i+10
	}
	fmt.Println(ma)

	fmt.Println(ma[0])
	ma[10]=10
	fmt.Println(ma)
	fmt.Println(ma[10])

	na := make(map[int]string)
	for i := 0; i <= 9; i++ {
		na[i] = strconv.Itoa(i+10)
	}
	fmt.Println(na)

	delete(na, 0)
	fmt.Println(na)
	fmt.Println(0)

	_, ok := na[1]  // called "comma ok syntax" will check to see if existence
	fmt.Println(ok) // can throw away "write only" variable
}
