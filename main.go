package main

import (
	"container/list"
	"fmt"
	"html/template"
	"net/http"
)

func signin(w http.ResponseWriter, r *http.Request) {
	var fileName = "signin.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error parsing file")
		return
	}
	type VetClinicName struct {
		Name string
	}
	err = t.ExecuteTemplate(w, fileName, VetClinicName{"Bobby's vet clinic"})
	if err != nil {
		fmt.Println("Error when executing template")
		return
	}
}

var userDB = map[string]string{
	"Roberta": "123456",
}

func signinSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Ok")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error")
	}
	fmt.Println(username, password)
}


func bookingWithTicket(w http.ResponseWriter, r *http.Request) {
	var queue = list.New()
	// Append
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	queue.PushBack(4)
	queue.PushBack(5)
	queue.PushBack(6)

	// Dequeue, remove from the front
	front := queue.Front()
	fmt.Println(front.Value)
	// Avoid memory leaks by removing the first element
	queue.Remove(front)

	var fileName = "booking.html"
	t, err := template.ParseFiles(fileName)
	err = t.ExecuteTemplate(w, fileName, queue)

	if err != nil {
		fmt.Println("Error when executing template booking")
		return
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello")
	case "/signin":
		signin(w, r)
	case "/signin-submit":
		signinSubmit(w, r)
	case "/booking":
		bookingWithTicket(w, r)
	}

	fmt.Printf("Handling function %s request\n", r.Method)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
	//Security - https
	//http.ListenAndServeTLS("", "cert.pem", "key.pem", nil)

}

//function with html as an example
/*func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello</h1>")
}*/
