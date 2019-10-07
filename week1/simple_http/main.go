package main

import (
	"fmt"
	"net/http"

	"github.com/thai2902/mono-go/week1/simple_http/handlers"
)

func main() {
	http.HandleFunc("/api/v1/students", handlers.GetStudents)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
	fmt.Println("Listening at 9090")
}
