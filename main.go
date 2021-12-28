package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()
	data := []byte(b)

	u := &User{}
	json.Unmarshal(data, u)
	resp, _ := json.Marshal(u)
	w.Write(resp)
	fmt.Printf("%s", b)

}
func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "./index.html")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	//http.Handle("/submit.html", http.FileServer(http.Dir("./")))
	http.Handle("/index.js", http.FileServer(http.Dir("./")))
	http.Handle("/eNWDJx.jpg", http.FileServer(http.Dir("./")))
	http.Handle("/style.css", http.FileServer(http.Dir("./")))
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
