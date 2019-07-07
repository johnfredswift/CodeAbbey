package InterviewCode

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// A single function from package http starts a web server.

/*
func LearnWebProgramming() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
*/


func LearnWebProgrammingTwo(){
	http.HandleFunc("/hello", HandleHello)
	fmt.Println("Serving on http://localhost:7777/hello")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))

}

func HandleHello(w http.ResponseWriter, req *http.Request){
	log.Println("serving", req.URL)
	fmt.Fprintln(w,"Welcome")
	currentTime := time.Now()
	fmt.Fprintln(w,"Current Date/Time is: " , currentTime.Format(time.RFC1123))
}


