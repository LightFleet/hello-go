package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	tryGo()
	

	fmt.Fprintf(w, "Hella World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)


	tryGo()
	tryGo()
	tryGo()
	tryGo()
	tryGo()
}

func tryGo() {
	var arr[3] int
	
	arr[0] = 42
	arr[1] = 24
	arr[2] = 01
	nums := [3]float64 {4.23, 5.23, 98.1}

	fmt.Println(arr[1])

	
	for i, value := range nums {
		fmt.Println(value, i)
	}
}