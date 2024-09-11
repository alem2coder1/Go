package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	// question 2
	item := 18
	if item > 0 {
		fmt.Println("positive")
	} else if item < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	weekDay := 3
	switch weekDay {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// In Go, the if statement allows variable declaration within the condition, and there are no parentheses required around the condition, unlike Java. Python doesn’t allow declarations in if conditions, and its syntax uses indentation instead of braces like Go and Java.

	arr := []int{1, 2, 3}
	for index, value := range arr {
		fmt.Println(index, value)
	}
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	a := 2
	b := 3
	result := add(a, b)
	fmt.Println("Sum:", result)

	c := "hello"
	d := "world"
	swappedA, swappedB := swap(c, d)
	fmt.Println("Swapped:", swappedA, swappedB)
	f := 10
	g := 5
	quotient, remainder := remainDer(f, g)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
	for _, value := range []int{1, 2, 3} {
		fmt.Println(value)
	}
	person := Person{Name: "alem", Age: 20}
	person.Greet()

	manager := Manager{
		Employee:   Employee{Name: "Alice", ID: 101},
		Department: "HR",
	}

	manager.Work()

	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// 调用 PrintArea 函数，分别传入 circle 和 rectangle
	PrintArea(circle)
	PrintArea(rectangle)

	var s Shape = Circle{Radius: 5}
	if c, ok := s.(Circle); ok {
		fmt.Println("s is a Circle with radius:", c.Radius)
	}

	switch shape := s.(type) {
	case Circle:
		fmt.Println("It's a Circle with radius:", shape.Radius)
	case Rectangle:
		fmt.Println("It's a Rectangle with width and height:", shape.Width, shape.Height)
	default:
		fmt.Println("Unknown shape")
	}

	product := Product{Name: "Laptop", Price: 999.99, Quantity: 10}

	jsonStr, err := ProductToJSON(product)
	if err != nil {
		fmt.Println("编码 JSON 时出错:", err)
		return
	}
	fmt.Println("JSON 输出:", jsonStr)

	decodedProduct, err := JSONToProduct(jsonStr)
	if err != nil {
		fmt.Println("解码 JSON 时出错:", err)
		return
	}
	fmt.Println("解码后的 Product:", decodedProduct)
}

func add(a int, b int) int {
	return a + b
}
func swap(c, d string) (string, string) {
	return d, c
}
func remainDer(f, g int) (int, int) {
	quotient := f / g
	remainder := f % g
	return quotient, remainder
}

// Exercise 4: 3
// func example(a, b int) (int, int) {
// 	return a + b, a * b
// }
// func divide(a, b int) (quotient, remainder int) {
// 	quotient = a / b
// 	remainder = a % b
// 	return
// }

// quotient, _ := divide(10, 3)
// OOP in Golang Exercise 1: Structs and Methods

type Person struct {
	Name string
	Age  int
}

// Method to greet for the Person struct
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Exercise 2: Embedding and Composition

type Employee struct {
	Name string
	ID   int
}

func (e Employee) Work() {
	fmt.Printf("Employee %s with ID %d is working.\n", e.Name, e.ID)
}

type Manager struct {
	Employee
	Department string
}

// 3
// 1. 定义接口 Shape
type Shape interface {
	Area() float64
}

// 2. Circle 结构体实现 Shape 接口
type Circle struct {
	Radius float64
}

// Circle 实现了 Area 方法，计算圆的面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 3. Rectangle 结构体实现 Shape 接口
type Rectangle struct {
	Width, Height float64
}

// Rectangle 实现了 Area 方法，计算矩形的面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 4. 定义一个函数 PrintArea，用来接收 Shape 接口并打印其面积
func PrintArea(s Shape) {
	fmt.Printf("The area is: %.2f\n", s.Area())
}

// 4
type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func ProductToJSON(p Product) (string, error) {

	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func JSONToProduct(jsonStr string) (Product, error) {
	var p Product
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		return p, err
	}
	return p, nil
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
if 20 > 18{
fmt.Println("20 is greater than 18")
}
func respective(item string, it int) {
	return
}
*/

//	for i := 0; i < 10; i++ {
//		fmt.Println(i)
//	}
