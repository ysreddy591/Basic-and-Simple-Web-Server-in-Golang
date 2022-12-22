package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Guys..This is a Simple basic website nad you cant interact with this")
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	/* err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
		return
	} */
	/* fname := r.FormValue("Fname")
	lname := r.FormValue("Lname")
	addr := r.FormValue("address") */
	fmt.Fprintf(w, "User Fullname is: %s\n", r.FormValue("Fname"))
	fmt.Fprintf(w, "User Lastname is: %s\n", r.FormValue("Lname"))
	fmt.Fprintf(w, "User Address is: %s\n", r.FormValue("address"))
}

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/form", formhandler)
	fmt.Println("Starting server at 6060....")
	log.Fatal(http.ListenAndServe(":6060", nil))

}
