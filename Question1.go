package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	x, _ := strconv.Atoi(os.Args[1]) //Converts the command line arg to an int

	p := fmt.Println //taken from gobyexample, makes a func variable to print

	now := time.Now()        //Time block taken from gobyexample, records current time
	var arr [10]int          //init array, cannot be variable
	for i := 0; i < x; i++ { //traverse array
		arr[i] = i
	}
	then := time.Now() //records end time
	p(now.Sub(then))   //prints the difference in two times

	now = time.Now()
	sli := make([]int, x) //create slice and traverse to assign values
	for i := 0; i < x; i++ {
		sli[i] = i
	}
	then = time.Now()
	p(now.Sub(then))

	now = time.Now()
	m := make(map[int]int) //Creates map
	for i := 0; i < x; i++ {
		m[i] = i
	}
	then = time.Now()
	p(now.Sub(then))

	//Increments each value in data structs by 1
	now = time.Now()
	for i := 0; i < x; i++ {
		arr[i] += 1
	}
	then = time.Now()
	p(now.Sub(then))

	now = time.Now()

	for i := 0; i < x; i++ {
		sli[i] += 1
	}
	then = time.Now()
	p(now.Sub(then))

	now = time.Now()
	for i := 0; i < x; i++ {
		m[i] += 1
	}
	then = time.Now()
	p(now.Sub(then))
}
