package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./index.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Fprintf(w, "activation link has been sent %v\n", r.PostForm)
		fmt.Printf("Post from website! r.PostFrom = %v\n", r.PostForm)
		Fname := r.FormValue("Fname")
		Lname := r.FormValue("Lname")
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Printf("First name = %s\n", Fname)
		fmt.Printf("Last Name = %s\n", Lname)
		fmt.Printf("Email = %s\n", email)
		fmt.Printf("password = %s\n", password)
		http.ServeFile(w, r, "./submit.html")
	default:
		fmt.Printf("Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.Handle("/index.js", http.FileServer(http.Dir("./")))
	http.Handle("/eNWDJx.jpg", http.FileServer(http.Dir("./")))
	http.Handle("/style.css", http.FileServer(http.Dir("./")))
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
