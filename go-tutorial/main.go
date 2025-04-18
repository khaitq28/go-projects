package main

import (
	"fmt"
	"github.com/google/uuid"
	"go-tutorial/internal/maths"
	"go-tutorial/internal/stringutils"
	"strconv"
	"strings"
)

func main() {

	println(stringutils.IsNumeric("1345012"))
	println(stringutils.IsNumeric("09b4t444"))
	fmt.Println(stringutils.CountTypes("eeeUUUU12"))
	println(strings.ToLower("Hello World"))

}

func testMain() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			chan1 <- i
		}
	}()
	go func() {
		for i := 100; i <= 105; i++ {
			chan2 <- i
		}
	}()
	//time.Sleep(time.Millisecond * 900)
	//for {
	//	select {
	//	case i := <-chan1:
	//		fmt.Println("chan1", i)
	//	case i := <-chan2:
	//		fmt.Println("chan2", i)
	//	default:
	//		fmt.Println("default")
	//		return
	//	}
	//}

	//argsWithProg := os.Args
	//fmt.Println("argsWithProg", argsWithProg)
	//fmt.Println("argsWithProg", os.Args[0])
	//fmt.Println("argsWithProg", os.Args[1])
	//fmt.Println("argsWithProg", os.Args[2])
	//fmt.Println("argsWithProg", os.Args[3])

	book := Book{name: "Goooo", author: "John"}
	fmt.Println("book", book.name, book.author)
	fmt.Println(stringutils.Concat("Hello", " vcl World"))
	fmt.Println(stringutils.Reverse("Hello World"))

	user := User{name: "John", age: uuid.New().ClockSequence(), email: "john@gmail.com"}
	user2 := User{name: "Johne", age: 12, email: "john@gmail.com"}
	user.Print()
	println("is equal : ", user.IsEqual(&user2))

	users := filterAdult([]User{user, user2})
	fmt.Println("Adult users:")
	fmt.Println(users)

	fmt.Println("run factorial of 10: ", maths.Factorial(10))

	s1 := "hello123"
	s2 := "321olleh"
	fmt.Println("is palindrome: ", s1, s2, stringutils.IsAnagram(s1, s2))
}

type Book struct {
	name   string
	author string
}

type Geo interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

func changeVal(i *int) {
	*i = 12
}

func (p *Person) changeName(x string) {
	p.name = x
}
func (p *Person) test() {
	fmt.Println("method", p.name, p.age)
	p.age = 12
	p.name = "Doe"
}

func test(p *Person) {
	p.name = "Change"
	fmt.Println("function ", p.name, p.age)
}

type Person struct {
	name string
	age  int
}

// a defer statement defers the execution of a function until the surrounding function returns
const (
	BIG   = 1 << 100
	SMALL = BIG >> 99
)

func change(i int) {
	i = 100
}

func loop() {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			println(strconv.Itoa(i) + " is odd")
		} else {
			println(strconv.Itoa(i) + " is even")
		}
	}
}
