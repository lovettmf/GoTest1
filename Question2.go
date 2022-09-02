package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	x, _ := strconv.Atoi(os.Args[1]) //Takes command line arg

	p := fmt.Println

	var arr [10]int
	for i := 0; i < x; i++ {
		rand.Seed(time.Now().UnixNano()) //new random number seed each iteration
		arr[i] = rand.Intn(100)          //assign random num
	}

	sli := make([]int, x)
	for i := 0; i < x; i++ {
		rand.Seed(time.Now().UnixNano())
		sli[i] = rand.Intn(100)
	}

	//sorting with sort and stable
	now := time.Now()
	sort.Ints(arr)
	then := time.Now()
	p(now.Sub(then))

	now = time.Now()
	sort.Ints(sli)
	then = time.Now()
	p(now.Sub(then))

}
