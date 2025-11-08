package helper

import "fmt"

// Here we are passing string, array of string and slice of string in function
// if you want to reurn slice of string just write []string

// We can use this variable across packages
var ExportedVar = "hello export"

// We can export function we just need matk first letter as capital
func Greet(name string, fruits [3]string, bookings []string) (string, []string) {
	fmt.Println(fruits)
	fmt.Println(bookings)
	return "Hello " + name, bookings
}
