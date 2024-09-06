package main

import "fmt"

func main() {
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	myslice2 := []int{1, 2, 3, 4, 5}
	myslice3 := append(myslice1, myslice2...)
	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("capacity = %d\n", cap(myslice1))
	fmt.Println(myslice3)
}

// var foo int
// var foo int = 42
// var foo, bar int = 42, 1302
// foo := 42\
/*var i int = 42
var f float64 = float64(i)
var u uint32 = uint32(f)
var a [10]int = [10]int{1, 2, 3, 4}
var (
     a int
     b int = 1
     c string = "hello"
   )
var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)
  HelloWorld
  fmt.Println(i,j)
  Hello World
 var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
  i has value: Hello and type: string
j has value: 15 and type: int
%v	Prints the value in the default format
%#v	Prints the value in Go-syntax format
%T	Prints the type of the value
%%	Prints the % sign

%b	Base 2
%d	Base 10
%+d	Base 10 and always show sign
%o	Base 8
%O	Base 8, with leading 0o
%x	Base 16, lowercase
%X	Base 16, uppercase
%#x	Base 16, with leading 0x
%4d	Pad with spaces (width 4, right justified)
%-4d	Pad with spaces (width 4, left justified)
%04d	Pad with zeroes (width 4
var array_name = [length]datatype{values}
arr2 := [5]int{4,5,6,7,8}
slice_name := make([]type, length, capacity)
const constant = "this is a constant"

func respective(item string, it int) {
	return
}
*/

//	for i := 0; i < 10; i++ {
//		fmt.Println(i)
//	}
