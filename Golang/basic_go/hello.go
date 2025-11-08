package main

import (
	"example/hello/helper"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// struct creation
type Udata struct {
	fname      string
	lname      string
	noofticket uint
}

func tempfunc(tempvar *string) {
	*tempvar = "updatedval"
}

func printNumbers() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		time.Sleep(1000 * time.Millisecond)
	}
	wg.Done()
}

// this is method
// here p is pass by value copy is used
// use pointer instead if required to change value
func (p Udata) printfname() {
	fmt.Println("Hello, my name is", p.fname)
}

func (p *Udata) Rename(newName string) {
	p.fname = newName
}

func fileoperations() {
	content := "This will be written in file"
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Print("length is ", length)
	databyte, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print("data is ", databyte)
}

func main() {

	wg.Add(1)
	defer fmt.Print("this will execute at last\n")
	defer fmt.Print("this will execute at last 2nd\n")
	// runs concurrently
	go printNumbers()

	// No need to specify type it can infer by itself
	var temp = "Hello world"
	const conftickets = 50
	var remainingTickets = conftickets
	fmt.Println(temp, "World!")
	fmt.Println("Total are ", conftickets, "Remaining are", remainingTickets)
	fmt.Printf("Total are %v Remaining are %v\n", conftickets, remainingTickets)

	fmt.Printf("Type is %T and another type is %T \n", conftickets, temp)

	// define type and assign values later
	var uname string
	var uticket int
	fmt.Scan(&uname)
	uticket = 20
	fmt.Println(uname, uticket)
	fmt.Printf("Address is %p", &uname)

	// assign and declare in same line
	decassign := 10
	fmt.Println(decassign)

	// Arrays assignment
	var prices = [3]int{10, 20, 30}
	fmt.Println(prices[0])
	fmt.Println(len(prices))

	arr1 := [5]int{1, 2}
	fmt.Println(arr1)
	fmt.Printf("Length of arr1 is %v\n", len(arr1))

	arr2 := [...]int{1, 2, 1, 1}
	fmt.Println(arr2)
	fmt.Printf("Length of arr2 is %v\n", len(arr2))
	fmt.Printf("Capacity of arr2 is %v\n", cap(arr2))

	// Here length is inferred based on values
	var array_name = [...]string{"a", "b"}
	fmt.Println(array_name)
	fmt.Println(len(array_name))
	fmt.Printf("Type of array is %T\n", array_name)

	// Slices in go that is dynamic arrays we dont need to specify length
	bookings := []string{}
	bookings = append(bookings, "Helllo abcddd")
	fmt.Printf("These are all our bookings: %v\n", bookings)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	space_str := "I love programming with Go!"
	fields := strings.Fields(space_str)
	for i, field := range fields {
		fmt.Println("Index: ", i)
		fmt.Println("Value: ", field)
	}

	// If else condition
	noTicketRemaining := remainingTickets == 0
	if noTicketRemaining {
		fmt.Println("Tickets are sold out")
	} else if remainingTickets == 1 {
		fmt.Println("1 Ticket remaining")
	} else {
		fmt.Println("More then 1 Tickets are available")
	}

	// Another way to write if else using condition only
	temp_index := 0
	for temp_index < 3 {
		fmt.Println(temp_index)
		temp_index++
	}

	and_condition := remainingTickets > 0 && remainingTickets < 5
	if and_condition {
		fmt.Println("More then 1 and less then 5 Tickets are available")
	}

	or_condition := remainingTickets == 0 || remainingTickets == 1
	if or_condition {
		fmt.Println("Either 0 or 1 ticket available")
	}

	// Infinite for loop
	// for {
	// 	fmt.Print("hello")
	// }

	switch uname {
	case "hello":
		fmt.Println("Hello")
	case "bye":
		fmt.Println("Bye")
	// consolidate multiple values into one case
	case "hello1", "hello2":
		fmt.Println("hello1 hello2")
	default:
		fmt.Println("Default switch printed ")
	}

	greeting, bookings := helper.Greet("abcd", fruits, bookings)
	fmt.Println(greeting)
	fmt.Println(bookings)
	tempvar := "hello"
	tempfunc(&tempvar)
	fmt.Println(tempvar)

	fmt.Println(helper.ExportedVar)

	// create an empty map
	var userData = make(map[string]string)
	userData["fname"] = "hello"
	userData["lname"] = "abcd"
	userData["tickets"] = strconv.FormatUint(uint64(10), 10)

	var emptyListMap = make([]map[string]string, 0)
	emptyListMap = append(emptyListMap, userData)

	fmt.Print(emptyListMap)

	for _, booking := range emptyListMap {
		fmt.Println(booking["fname"])
	}
	// tempmap.append("hello", "abcd")

	var ulst = make([]Udata, 0)
	ulst = append(ulst, Udata{
		fname:      "hello",
		lname:      "abcd",
		noofticket: 10,
	})

	var udata2 = Udata{"hellodata2", "abcd2", 20}
	fmt.Println(udata2)
	udata2.printfname()

	fmt.Println(ulst)
	for _, booking := range ulst {
		fmt.Println(booking.fname)
	}

	//Pointers
	var tempptr = &tempvar
	fmt.Printf("pointer val is %v\n", *tempptr)
	*tempptr = "newtempptr"
	fmt.Printf("updated pointer val is %v\n", *tempptr)

	fileoperations()
	// wait until it becomes 0
	wg.Wait()
}
