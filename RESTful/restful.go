package main

import (
	"fmt"
	"net/http"
)

//  http://localhost:8080/users

func main() {

	http.HandleFunc("/users", handleJosnUsers)

	http.ListenAndServe(":8080", nil)

}

func handleJosnUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "ID:1,Name:张三")

	fmt.Fprintln(w, "ID:2,Name:李四")

	fmt.Fprintln(w, "ID:3,Name:王五")

}
