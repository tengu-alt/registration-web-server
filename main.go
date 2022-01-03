package main

import (
	"encoding/json"
	"fmt"
	"functions"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}
type ValidationErr struct {
	FieldValue string
	ErrMassage string
}

func (u *User) Validate() []ValidationErr {
	errors := make([]ValidationErr, 0, 0)
	if len(u.FirstName) < 2 {
		errors = append(errors, ValidationErr{
			FieldValue: "FirstName",
			ErrMassage: fmt.Sprintf("field %s length should be equal or longer than 2", "FirstName"),
		})
	}
	if len(u.FirstName) > 64 {
		errors = append(errors, ValidationErr{
			FieldValue: "FirstName",
			ErrMassage: fmt.Sprintf("field %s length should be less than 64", "FirstName"),
		})
	}

	if len(u.LastName) < 2 {
		errors = append(errors, ValidationErr{
			FieldValue: "LastName",
			ErrMassage: fmt.Sprintf("field %s length should be equal or longer than 2", "LastName"),
		})
	}

	if len(u.LastName) > 64 {
		errors = append(errors, ValidationErr{
			FieldValue: "LastName",
			ErrMassage: fmt.Sprintf("field %s length should be less than 64", "LastName"),
		})
	}
	if len(u.Password) < 8 {
		errors = append(errors, ValidationErr{
			FieldValue: "Password",
			ErrMassage: fmt.Sprintf("field %s length should be equal or longer than 8", "Password"),
		})
	}

	if len(u.Password) > 64 {
		errors = append(errors, ValidationErr{
			FieldValue: "Password",
			ErrMassage: fmt.Sprintf("field %s length should be less than 64", "Password"),
		})
	}

	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(u.Email) != true {

		errors = append(errors, ValidationErr{
			FieldValue: "Email",
			ErrMassage: "email failed verification",
		})
	}

	return errors
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
	validationErrors := u.Validate()
	fmt.Printf("%s", validationErrors)
	if len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)

		b, err := json.Marshal(&validationErrors)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
		return
	}
	w.Write([]byte("[{}]"))
	fmt.Printf(functions.ValiderEmail())
	//resp, _ := json.Marshal(u)
	//w.Write()
	//fmt.Printf("%s", b)

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
	http.Handle("/submit.html", http.FileServer(http.Dir("./")))
	http.Handle("/index.js", http.FileServer(http.Dir("./")))
	http.Handle("/eNWDJx.jpg", http.FileServer(http.Dir("./")))
	http.Handle("/style.css", http.FileServer(http.Dir("./")))
	fmt.Printf("Starting server for testing HTTP POST in 8081...\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
